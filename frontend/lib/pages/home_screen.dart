import 'package:flutter/material.dart';
import 'package:frontend/pages/login_screen.dart';
import 'package:shared_preferences/shared_preferences.dart';

class HomeScreen extends StatefulWidget {
  final SharedPreferences preferences;
  const HomeScreen({super.key, required this.preferences});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  String? session = '';
  @override
  Widget build(BuildContext context) {
    final session = widget.preferences.getString("session");

    return Scaffold(
      body: Center(
        child: Padding(
          padding: const EdgeInsets.all(10),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Text("Your Session is $session"),
              ElevatedButton(
                onPressed: () {
                  widget.preferences.clear();
                  Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (context) => LoginScreen(
                        preferences: widget.preferences,
                      ),
                    ),
                  );
                },
                child: const Text("Log Out!"),
              )
            ],
          ),
        ),
      ),
    );
  }
}
