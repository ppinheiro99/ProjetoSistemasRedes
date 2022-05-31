import 'dart:async';
import 'dart:convert';
import 'package:bloc_login/model/user_model.dart';
import 'package:bloc_login/service/global_variable.dart';
import 'package:flutter/cupertino.dart';
import 'package:http/http.dart' as http;
import 'package:bloc_login/model/api_model.dart';
import 'package:bloc_login/model/route_model.dart';
import 'dart:developer' as developer;
import 'package:bloc_login/dao/user_dao.dart';

// server
final _base = "http://18.130.231.194:8080/api/mobile/";
// localhost
//final _base = "http://10.0.2.2:8080/api/";
//

final _tokenEndpoint = "mobileLogin";
final _tokenURL = _base + _tokenEndpoint;

final _getRoute = _base + "getDriverRoutes/";
final _finishedRoute = _base + "finishedRoute/";
int idTruckDriver = 0;

Future<Token> getToken(UserLogin userLogin) async {
  print(_tokenURL);
  print(userLogin);
  final http.Response response = await http.post(
    _tokenURL,
    headers: <String, String>{
      'Content-Type': 'application/json; charset=UTF-8',
    },
    body: jsonEncode(userLogin.toDatabaseJson()),
  );

  if (response.statusCode == 200) {
    return Token.fromJson(json.decode(response.body));
  } else {
    print(json.decode(response.body).toString());
    throw Exception(json.decode(response.body));
  }
}

Future<RouteModel> getRoute(String username) async {
  print(_getRoute + username);
  print(username);
  final http.Response response = await http.get(
    _getRoute + username,
    headers: <String, String>{
      'Content-Type': 'application/json;',
    },
  );

  if (response.statusCode == 200) {
    return RouteModel.fromJson(json.decode(response.body));
  } else {
    print(json.decode(response.body).toString());
    throw Exception(json.decode(response.body));
  }
}

Future<RouteModel> finishedRoute(int idDisplacement) async {
  if (idDisplacement != 0) {
    print(_finishedRoute + (idDisplacement.toString()));
    final http.Response response = await http.get(
      _finishedRoute + (idDisplacement.toString()),
      headers: <String, String>{
        'Content-Type': 'application/json;',
      },
    );
    if (response.statusCode == 200) {
      print(json.decode(response.body));
    } else {
      print(json.decode(response.body).toString());
      throw Exception(json.decode(response.body));
    }
  }
}
