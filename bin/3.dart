import 'dart:math';

void main(List<String> args) {
  var strs = ["abcabcbb", "bbbbb", "pwwkew"];
  for (var str in strs) {
    print(Solution.lengthOfLongestSubstring(str));
  }
}

class Solution {
  static int lengthOfLongestSubstring(String str) {
    int l = 0;
    int r = 0;
    // 初始化Set
    Set<String> set = {};

    int res = 0;

    while (r < str.length) {
      while (set.contains(str[r])) {
        set.remove(str[l]);
        l++;
      }
      set.add(str[r]);
      res = max(res, set.length);
      r++;
    }

    return res;
  }
}
