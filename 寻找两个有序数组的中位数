
//找到两个有序数组中的第K个数思路
class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        const int m = nums1.size();
        const int n = nums2.size();
        int total = m + n;
        if(total % 2 == 1){//说明是奇数个
            return findkth(nums1,nums2,0,0,total/2+1);
        }else{
            return (findkth(nums1,nums2,0,0,total/2)+findkth(nums1,nums2,0,0,total/2+1))/2.0;
        }
    }
    
    int findkth(vector<int>& nums1,vector<int>& nums2,int m,int n,int k){
        int res=0,cnt=0;
        while(m < nums1.size() && n < nums2.size() && cnt < k){
            if(nums1[m] < nums2[n]){
                res = nums1[m];
                m++;
                cnt++;
            }else{
                res = nums2[n];
                n++;
                cnt++;
            }
        }
        
        while(m < nums1.size() && cnt < k){
            res = nums1[m];
            m++;
            cnt++;
        }
        while(n < nums2.size() && cnt < k){
            res = nums2[n];
            n++;
            cnt++;
        }
        return res;
    }
};
