import 'dart:convert';

class User {
  final String user;

  User({this.user});

  factory User.fromJson(dynamic json) {
    return User(
      user: json['user'],
    );
  }
}

main() {
  const data = '[{"user": "jy"}, {"user":"kk"}]';

  var parsed = jsonDecode(data);
  List<User> users = List<User>.from(parsed.map((r) => User.fromJson(r)));

  print(users[0].user);
}
