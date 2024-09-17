import 'package:shared_preferences/shared_preferences.dart';

enum AuthState { token, noToken, registered }

class SessionController {
  final SharedPreferences preferences;

  SessionController({required this.preferences});

  Future<SharedPreferences> getSharedPreferences() async {
    return await SharedPreferences.getInstance();
  }

  Future<String?> getSession() async {
    return preferences.getString("session");
  }

  void createSession(String session) async {
    print("Creating Session!");
    preferences.setString("session", session);
  }
}
