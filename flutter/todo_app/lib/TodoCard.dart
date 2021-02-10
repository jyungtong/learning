import 'package:flutter/material.dart';

class TodoCard extends StatefulWidget {
  String task;
  bool isCompleted;

  TodoCard({this.task, this.isCompleted = false});

  @override
  _TodoCardState createState() => _TodoCardState();
}

class _TodoCardState extends State<TodoCard> {
  final _textController = TextEditingController();
  FocusNode _focus = new FocusNode();

  @override
  void initState() {
    super.initState();
    _focus.addListener(_onFocusChange);
    _focus.requestFocus();
  }

  void _onFocusChange() {
    setState(() {
      widget.task = _textController.text;
    });
  }

  @override
  Widget build(BuildContext context) {
    final isTaskEmpty = ['', null].contains(widget.task);

    return Container(
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Row(
          children: [
            Checkbox(
              onChanged: (bool value) {
                setState(() {
                  widget.isCompleted = value;
                });
              },
              value: widget.isCompleted,
            ),
            Expanded(
              child: isTaskEmpty
                  ? TextField(
                      focusNode: _focus,
                      controller: _textController,
                      decoration: InputDecoration(
                        hintText: 'Type something',
                      ),
                      style: TextStyle(
                        fontSize: 26,
                      ),
                    )
                  : Text(
                      widget.task,
                      style: widget.isCompleted
                          ? TextStyle(
                              decoration: TextDecoration.lineThrough,
                              fontStyle: FontStyle.italic,
                              fontSize: 26,
                            )
                          : TextStyle(fontSize: 26),
                    ),
            )
          ],
        ),
      ),
    );
  }
}
