class Solution {
public:
    int reverse(int x) {
        long long anwser=0;
        while(x!=0)
        {
            anwser=anwser*10+(x%10);
            x/=10;
        }
        if(anwser<INT_MIN || anwser>INT_MAX)
        {
            anwser=0;
        }
        return anwser;
    }
};
