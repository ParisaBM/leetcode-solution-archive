class Solution:
    def candy(self, ratings) -> int:
        amount = [1] * len(ratings)
        ordered = sorted(enumerate(ratings), key = lambda x: x[1])
        for x in ordered:
            if x[0] > 0 and ratings[x[0]] > ratings[x[0]-1]:
                amount[x[0]] = amount[x[0]-1] + 1
            if x[0] < len(ratings)-1 and ratings[x[0]] > ratings[x[0]+1]:
                amount[x[0]] = max(amount[x[0]], amount[x[0]+1] + 1)
        return sum(amount)

s = Solution()
print(s.candy([1,0,2]))
print(s.candy([1,2,2]))
