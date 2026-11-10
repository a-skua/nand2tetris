import 'package:test/test.dart';
import 'package:assembler/parser.dart';

void main() {
  group('Parser', () {
    test('should parse a simple instruction', () {
      final parser = Parser(
        [
          // A-instruction
          '@1',
          // C-instruction
          'D=A',
        ].join('\n'),
      );

      expect(parser.hasMoreLines, isTrue);
      parser.advance();
      expect(parser.instructionType, InstructionType.a);

      expect(parser.hasMoreLines, isTrue);
      parser.advance();
      expect(parser.instructionType, InstructionType.c);

      // No more lines
      expect(parser.hasMoreLines, isFalse);
    });
  });
}
