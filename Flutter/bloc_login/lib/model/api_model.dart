class UserLogin {
  String username;
  String password;
  int idCamionista;

  UserLogin({this.username, this.password, this.idCamionista});

  Map <String, dynamic> toDatabaseJson() => {
    "email": this.username,
    "password": this.password
  };
}

class Token{
  String token;
  int idCamionista;
  Token({this.token, this.idCamionista});

  factory Token.fromJson(Map<String, dynamic> json) {
    return Token(
      token: json['token'],
      idCamionista: json['ID']
    );
  }
}

