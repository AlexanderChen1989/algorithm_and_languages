import 'dart:math';

void main(List<String> args) {
  print(characterReplacement('ABAB', 2));
  print(characterReplacement('AABABBA', 1));
}

int characterReplacement(String s, int k) {
  var res = 0;

  // 初始化数组是方式
  var counter = List.generate(26, (index) => 0);

  var l = 0;
  // dart的字符串是utf16
  var a = 'A'.codeUnitAt(0);
  for (var r = 0; r < s.length; r++) {
    var idx = s.codeUnitAt(r) - a;
    counter[idx]++;
    var len = r - l + 1;
    if (len - maxCounter(counter) <= k) {
      res = max(res, len);
    }
    while (r - l + 1 - maxCounter(counter) > k) {
      var idx = s.codeUnitAt(l) - a;
      counter[idx]--;
      l++;
    }
  }

  return res;
}

int maxCounter(List<int> counter) {
  var res = counter[0];
  for (var v in counter) {
    if (res < v) {
      res = v;
    }
  }
  return res;
}
