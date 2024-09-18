import 'package:frontend/controller/session_controller.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';

class AuthController {
  static const IPADDRESS = "192.168.103.88";
  Future<int> registerUser(String name, String email, String password) async {
    try {
      final postRequest = await http.post(
        Uri.parse("http://$IPADDRESS:1337/api/v1/auth/register"),
        headers: <String, String>{'Content-Type': 'application/json'},
        body: jsonEncode(
          <String, String>{
            'username': email,
            'name': name,
            'password': password,
          },
        ),
      );
      print("Post Request Response!");
      print(postRequest.statusCode);
      switch (postRequest.statusCode) {
        case 200:
          final response = jsonDecode(postRequest.body);
          print('obtained:)');
          print(response);
          print(response.accessToken);
          final session = response["accessToken"];
          final sharedPref = await SharedPreferences.getInstance();
          final sessionController = SessionController(preferences: sharedPref);
          sessionController.createSession(session);
          break;
        case 400:
          return 404;
        // return 404
        default:
          return 500;
        // return 500
      }
      print(postRequest.body);
      return postRequest.statusCode;
    } catch (e) {
      print("Error Hit");
      print(e);
      return 1;
    }
  }
}
//   void parseToken(String responseBody) {
//     final parsed =
//         (jsonDecode(responseBody) as List).cast<Map<String, dynamic>>();

//     print(parsed);
//     print(parsed["accessToken"]);
//   }
// }
