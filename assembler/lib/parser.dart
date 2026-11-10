class Parser {
  final String _code;
  int _index = -1;
  int _lastIndex = 0;
  Parser(this._code);

  String get _line => _code.substring(_index, _lastIndex).trim();

  bool get hasMoreLines => _lastIndex < _code.length;

  void advance() {
    if (hasMoreLines) {
      _index = _lastIndex;
      _lastIndex += switch (_code.indexOf('\n', _index)) {
        -1 => _code.length - _index,
        final index => index + 1,
      };
    }
  }

  InstructionType get instructionType => switch (_line.substring(0, 1)) {
    '@' => InstructionType.a,
    '(' => InstructionType.l,
    _ => InstructionType.c,
  };
}

enum InstructionType { a, c, l }
