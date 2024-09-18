import 'package:flutter/material.dart';
import 'package:frontend/controller/session_controller.dart';
import 'package:frontend/pages/home_screen.dart';
import 'package:frontend/pages/login_screen.dart';
import 'package:shared_preferences/shared_preferences.dart';

class InitialPage extends StatefulWidget {
  final SharedPreferences prefs;

  const InitialPage({super.key, required this.prefs});

  @override
  State<InitialPage> createState() => _InitialPageState();
}

class _InitialPageState extends State<InitialPage> {
  @override
  Widget build(BuildContext context) {
    String? token = widget.prefs.getString("session");
    AuthState state = AuthState.noToken;
    if (token != null) {
      state = AuthState.token;
    }
    switch (state) {
      case AuthState.token:
        return HomeScreen(
          preferences: widget.prefs,
        );
      case AuthState.noToken:
        print("login screen!");
        return LoginScreen(preferences: widget.prefs);
      default:
    }
    return const Placeholder();
  }
}
