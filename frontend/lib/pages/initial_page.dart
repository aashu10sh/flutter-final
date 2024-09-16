import 'package:flutter/material.dart';
import 'package:frontend/controller/page_controller.dart';
import 'package:frontend/pages/login_screen.dart';

class InitialPage extends StatefulWidget {
  const InitialPage({super.key});

  @override
  State<InitialPage> createState() => _InitialPageState();
}

class _InitialPageState extends State<InitialPage> {
  late AuthState state;

  @override
  Widget build(BuildContext context) {
    @override
    void initState() {
      super.initState();
      state = AuthState.noToken;
    }

    switch (state) {
      case AuthState.token:
        break;

      case AuthState.noToken:
        return LoginScreen();
      default:
    }
    return const Placeholder();
  }
}
