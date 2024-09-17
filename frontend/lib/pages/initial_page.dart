import 'package:flutter/material.dart';
import 'package:frontend/controller/session_controller.dart';
import 'package:frontend/pages/home_screen.dart';
import 'package:frontend/pages/login_screen.dart';
import 'package:shared_preferences/shared_preferences.dart';

class InitialPage extends StatefulWidget {
  const InitialPage({super.key});

  @override
  State<InitialPage> createState() => _InitialPageState();
}

class _InitialPageState extends State<InitialPage> {
  late AuthState state = AuthState.noToken;

  @override
  Widget build(BuildContext context) {
    @override
    void initState() async {
      super.initState();
      print("initial state runs!");
      final sharedPreference = await SharedPreferences.getInstance();
      final SessionController sessionController =
          SessionController(preferences: sharedPreference);
      final session = await sessionController.getSession();
      if (session == null) {
        state = AuthState.noToken;
      } else {
        state = AuthState.token;
      }
    }

    switch (state) {
      case AuthState.token:
        return HomeScreen();
      case AuthState.noToken:
        print("login screen!");
        return LoginScreen();
      default:
    }
    return const Placeholder();
  }
}
