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
      theme: ThemeData(primarySwatch: Colors.deepOrange),
      home: InitialPage(),
    );
  }
}
