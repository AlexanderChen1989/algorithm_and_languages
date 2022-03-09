#include <iostream>
#include <string>
#include <tuple>

class Solution
{
public:
    int characterReplacement(std::string s, int k)
    {
        int counter[26] = {0};

        int res = 0;
        int l = 0;
        for (int r = 0; r < s.length(); r++)
        {
            counter[s[r] - 'A']++;
            int len = r - l + 1;
            if (len - maxCounter(counter) <= k)
            {
                if (res < len)
                {
                    res = len;
                }
            }
            while (r - l + 1 - maxCounter(counter) > k)
            {
                counter[s[l] - 'A']--;
                l++;
            }
        }

        return res;
    }

    int maxCounter(const int counter[])
    {
        int res = counter[0];
        for (int i = 0; i < 26; i++)
        {
            if (res < counter[i])
            {
                res = counter[i];
            }
        }
        return res;
    }
};

int main()
{
    std::pair<std::string, int> tests[] = {{"ABAB", 2}, {"AABABBA", 1}};

    Solution solution;
    for (auto test : tests)
    {
        std::string s;
        int k;
        std::tie(s, k) = test;
        std::cout << solution.characterReplacement(s, k) << std::endl;
    }

    return 0;
}
