void main(List<String> args) {
  var s = "a", t = "aa";

  print(Solution.minWindow(s, t));
}

class Solution {
  static String minWindow(String s, String t) {
    if (t == "" || s.length < t.length) return "";
    var counterM = <String, int>{}, windowM = <String, int>{};

    for (int i = 0; i < t.length; i++) {
      counterM.update(t[i], (value) => value += 1, ifAbsent: () => 1);
    }

    var need = counterM.length, has = 0, l = 0, r = 0;
    var res = [-1, -1], resLen = 10 ^ 7;

    for (; r < s.length; r++) {
      var c = s[r];
      windowM.update(c, (value) => value + 1, ifAbsent: () => 1);
      if (counterM.containsKey(c) && counterM[c] == windowM[c]) {
        has += 1;
      }
      while (has == need) {
        var tLen = r - l + 1;
        if (tLen < resLen) {
          res = [l, r];
          resLen = tLen;
        }
        var c = s[l];
        windowM[c] = windowM[c]! - 1;
        if (counterM.containsKey(c) && windowM[c]! < counterM[c]!) {
          has -= 1;
        }
        l += 1;
      }
    }

    return s.substring(res[0], res[1] + 1);
  }
}
