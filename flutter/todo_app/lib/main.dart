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
          child: Scrollbar(
            child: ListView.builder(
              itemCount: _todoCards.length,
              itemBuilder: (BuildContext context, int index) {
                return Dismissible(
                  key: UniqueKey(),
                  background: Container(
                    color: Colors.red,
                    alignment: Alignment.centerRight,
                    padding: EdgeInsets.all(8.0),
                    child: Text(
                      'Delete',
                      style: TextStyle(color: Colors.white),
                    ),
                  ),
                  child: _todoCards[index],
                  onDismissed: (direction) {
                    print('====direction: $direction');

                    setState(() {
                      _todoCards.removeAt(index);
                    });
                  },
                );
              },
            ),
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
