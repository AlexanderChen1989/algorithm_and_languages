#include <iostream>
#include <tuple>
#include <set>

class Solution
{
public:
	int lengthOfLongestSubstring(std::string s)
	{
		std::set<char> counter;

		int res = 0, l = 0;
		for (int r = 0; r < s.length(); r++)
		{
			while (counter.find(s[r]) != counter.end())
			{
				counter.erase(s[l]);
				l++;
			}
			counter.insert(s[r]);
			if (res < counter.size())
			{
				res = counter.size();
			}
		}

		return res;
	}
};

int main()
{
	std::string ss[] = {"abcabcbb", "bbbbb", "pwwkew"};
	Solution solution;
	for (auto s : ss)
	{
		std::cout << solution.lengthOfLongestSubstring(s) << std::endl;
	}
	return 0;
}