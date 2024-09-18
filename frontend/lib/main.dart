import 'package:flutter/material.dart';
import 'package:frontend/pages/initial_page.dart';
import 'package:shared_preferences/shared_preferences.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  final SharedPreferences prefs = await SharedPreferences.getInstance();
  runApp(MainApplication(
    preferences: prefs,
  ));
}

class MainApplication extends StatelessWidget {
  final SharedPreferences preferences;
  const MainApplication({super.key, required this.preferences});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        primarySwatch: Colors.green,
        textTheme: const TextTheme(
            labelSmall: TextStyle(color: Colors.white),
            labelMedium: TextStyle(color: Colors.white),
            labelLarge: TextStyle(color: Colors.white)),
      ),
      home: InitialPage(
        prefs: preferences,
      ),
    );
  }
}
