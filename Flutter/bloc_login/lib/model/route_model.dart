import 'dart:developer' as developer;

class RouteModel{
  String coords;
  int data;
  RouteModel({this.coords, this.data});

  factory RouteModel.fromJson(Map<String, dynamic> json) {
    return RouteModel(
        coords: json['coords'],
        data: json['data']
    );
  }
}