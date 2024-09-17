import 'package:flutter/material.dart';
import 'package:frontend/pages/initial_page.dart';

void main() {
  runApp(MainApplication());
}

class MainApplication extends StatelessWidget {
  const MainApplication({super.key});

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
      home: InitialPage(),
    );
  }
}
