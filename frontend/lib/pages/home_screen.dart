import 'package:flutter/material.dart';
import 'package:frontend/controller/session_controller.dart';
import 'package:shared_preferences/shared_preferences.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  String? session = '';
  @override
  Widget build(BuildContext context) {
    @override
    void initState() async {
      super.initState();
      SharedPreferences sharedpref = await SharedPreferences.getInstance();
      SessionController sessionController =
          SessionController(preferences: sharedpref);
      setState(() async {
        session = await sharedpref.getString("session");
        print("session obtained is $session");
      });
    }

    return Scaffold(
      body: Center(
        child: Text("Your Session is $session"),
      ),
    );
  }
}
