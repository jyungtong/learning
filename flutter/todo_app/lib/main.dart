import 'package:flutter/material.dart';
import './TodoCard.dart';

void main() {
  runApp(TodoApp());
}

class TodoApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Todo App',
      home: TodoScreen(),
      theme: ThemeData(
        textTheme: TextTheme(
          headline6: TextStyle(fontSize: 26.0, fontStyle: FontStyle.italic),
        ),
      ),
    );
  }
}

class TodoScreen extends StatefulWidget {
  @override
  _TodoScreenState createState() => _TodoScreenState();
}

class _TodoScreenState extends State<TodoScreen> {
  final List<TodoCard> _todoCards = [
    // TodoCard(task: 'Task 1'),
    // TodoCard(task: 'Task 2', isCompleted: true),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Todo'),
      ),
      body: GestureDetector(
        onTap: () {
          FocusScope.of(context).unfocus();
        },
        child: Container(
          child: ListView.builder(
            itemBuilder: (BuildContext context, int index) => _todoCards[index],
            itemCount: _todoCards.length,
          ),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        child: Icon(Icons.add),
        onPressed: () {
          setState(() {
            _todoCards.add(TodoCard());
          });
        },
      ),
    );
  }
}
