class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        unordered_map<char,int> umap;
        int j =0; 
        int length =0;
        for(int i=0;i<s.size();i++){
            if(umap.count(s[i])){
                j = max(j,umap[s[i]]+1);
            }
            umap[s[i]] = i;
           length = max(length,i-j+1);
        }
        return length;
    }
};
