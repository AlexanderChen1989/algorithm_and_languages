void main(List<String> arguments) {
  // 类型推导、变量定义、变量初始化、数组

  // var prices;
  // prices = [7, 1, 5, 3, 6, 4];

  // List<int> prices;
  // prices = [7, 1, 5, 3, 6, 4];

  var prices = [7, 1, 5, 3, 6, 4];

  // var prices = <int>[7, 1, 5, 3, 6, 4];
  // List<int> prices = <int>[7, 1, 5, 3, 6, 4];


  print(Solution.maxProfitClass(prices));

  // Solution实例化
  var solution = Solution();
  // solution = new Solution(); // new可选的
  print(solution.maxProfitInstance(prices));

  solution = Solution.withPrices(prices);
  // solution = new Solution.withPrices(prices); // new可选
  print(solution.maxProfitInstance(null));
}

class Hello {
  final String name;
  final int age;
  // 普通构造函数
  Hello()
      : name = '',
        age = 10;
  // const 构造函数需要全部的字段都是final的
  const Hello.withName(this.name) : age = 100;
}

class Solution {
  List<int> prices;

  // 默认构造函数
  Solution() : prices = [];

  // 其他构造函数
  Solution.withPrices(this.prices);

  // 实例方法
  int maxProfitInstance(List<int>? prices_) {
    prices_ = prices_ ?? prices;

    return Solution.maxProfitClass(prices_);
  }

  // 类方法
  static int maxProfitClass(List<int> prices) {
    int l = 0;
    int r = 0;
    int max = 0;

    while (l < prices.length && r < prices.length) {
      var diff = prices[r] - prices[l];
      if (diff <= 0) {
        l = r;
        r = l + 1;
        continue;
      }
      if (diff > max) {
        max = diff;
      }
      r = r + 1;
    }

    return max;
  }
}
