import 'dart:async';

import 'package:bloc_login/api_connection/api_connection.dart';
import 'package:bloc_login/model/route_model.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:bloc_login/bloc/authentication_bloc.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:geolocator/geolocator.dart';
import 'package:latlong/latlong.dart';
import 'package:bloc_login/dao/user_dao.dart';

   //String route;
  var points = <LatLng>[
  ];
  RouteModel route;

  var locationMessage= "";
  final userDao = UserDao();
  var position;
  var _latitude = 41.17323531013999;
  var _longitude = -8.611167669296265;
  bool getroute = false;
  Marker marker;
  List<Marker> markers;
  Timer _locationTimer;

class HomePageScreen extends StatefulWidget {

  HomePageScreen();

  @override
  HomePageScreenState createState() {
    return HomePageScreenState();
  }
}

class HomePageScreenState extends State<HomePageScreen> {

  static final MapOptions initialLocation = MapOptions(
    center: LatLng(_latitude, _longitude),
    zoom: 18.0,
  );

  MapController _controller = MapController();

  void getCoords() async {
    position = await Geolocator().getCurrentPosition(desiredAccuracy: LocationAccuracy.high); // para pegar na localização atual do camionista
    if(_controller != null){
       // colocar o camião posicionado na Posição correta
       _controller.center.latitude = position.latitude;
       _controller.center.longitude = position.longitude;
       // Update Truck
         markers[0] = Marker(
             width: 45.0,
             height: 45.0,
             point: _controller.center,
             anchorPos: AnchorPos.align(AnchorAlign.center),
             builder: (context) => Container(
               child: Icon(
                 Icons.adjust_outlined,
                 size: 45,
                 color: Colors.blue
               ),
             )
         );
       _controller.moveAndRotate(_controller.center, 18.0, 30); // Novas coordenadas, zoom e o ângulo pretendido
     } // assim que as coordenadas do camiao forem iguais às coordenadas finais da rota, temos de enviar o id da rota para o backEnd
    if(getroute == true){
      var latTruck = _controller.center.latitude.toString().substring(0,5);
      var longTruck = _controller.center.longitude.toString().substring(0,5);
      var latFinished = points.last.latitude.toString().substring(0,5);
      var longFinished = points.last.longitude.toString().substring(0,5);

      // assim que a localização atual do camionista for igual às coordenadas finais do deslocamento, temos de enviar para o golang
      // Basta os primeiros 2 dígitos da latitude e os 2 digitos da longitude
      // (para do lado do golang passar o deslocamento para o mapa de viagem (mapa de viagem é um conjunto de deslocamentos concluidos pelos camionistas)
      // e eliminar o deslocamento da bd de deslocamentos )
      if(latTruck == latFinished && longTruck == longFinished){
        finishedRoute(route.data);
        getroute = false; // como ja nao existe rota, colocamos a false para nao entrar na condicao de cima
        points = []; // limpamos o array de points
      }
    }
  }

  void drawRoutes() async {
    List<Map> u =await userDao.getUser(0).then((value) {
      return value;
    });
    var username;
    u[0].forEach((key, value) {
      if(key == "username"){
        username = value;
      }
    });

    route = await getRoute(username).then((logged) { // Vai buscar as coordenadas que o user recebeu do pedido à BD
      return logged;
    });
    if(route.coords != "") { // caso o user ainda nao tenha rotas criadas pelo chefe de tráfego
      getroute = true;
      print(route.data);
      var delete1 = route.coords.replaceAll(RegExp('[[]'), ''); // para retirar o [ da string
      var delete2 = delete1.replaceAll(RegExp(']'), ''); // para retirar o ] da string
      List result = delete2.split(','); // para converter a String em Array ( até chegar à , é uma posição )
      for (var i = 0; i <= result.length - 1; i += 2) {
        points.add(new LatLng(double.parse(result[i]), double.parse(result[i + 1])));
      }
    }
  }

  @override
  void initState() {
    super.initState();
    markers = [Marker()];
    drawRoutes();
    //getCoords();
    const oneSec = const Duration(seconds:1);
    _locationTimer = Timer.periodic(oneSec, (Timer t) =>getCoords()); // para pegar as coordenadas recebidas no telemóvel
  }

  @override
  void dispose() { // Quebra o "ciclo" do timer
    _locationTimer.cancel();
    markers.clear();
    getroute = false;
    points = []; // limpamos o array de points
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Mapa'),
        actions: <Widget>[
          new IconButton(
            icon: new Icon(Icons.login),
            onPressed: () {
              dispose();
              BlocProvider.of<AuthenticationBloc>(context)
                  .add(LoggedOut());
            },
          )
        ],
      ),
      body:
         FlutterMap(
          mapController: _controller,
          options: initialLocation,
          layers: [
            new TileLayerOptions(
              urlTemplate:"https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
                subdomains: ['a', 'b', 'c'],
            ),
            MarkerLayerOptions(markers: markers),
            new PolylineLayerOptions(
              polylines: [
                new Polyline(
                  points: points,
                  strokeWidth: 5.0,
                  color: Colors.blue
                )
              ]
            )
          ],
      ),
    );
  }
}
