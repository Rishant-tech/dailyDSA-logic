#include<bits/stdc++.h>
using namespace std;
vector<int> solve(vector<int> &nums,int k){
        int i=0;
        int n=nums.size();
        vector<int> maxi;
        deque<int> dq;
        for(int j=0;j<nums.size();j++){
            while(!dq.empty() && nums[dq.back()]<nums[j]){
                dq.pop_back();
            }
            dq.push_back(j);
            if(j-i+1 > k){
               if(dq.front()==i){
                 dq.pop_front();
               }
               i++;
            }
            if(j-i+1 == k){
                maxi.push_back(nums[dq.front()]);
            }
        }
        return maxi;
}
int main(){
    int n;
    cout<<"enter size of an array";
    cin>>n;
    vector<int> arr(n);
    cout<<"enter the array";
    for(int i=0;i<n;i++){
        cin>>arr[i];
    }
    int k;
    cout<<"enter k";
    cin>>k;
    vector<int> res = solve(arr,k);
    for(int r:res){
        cout<<r<<" ";
    }
    return 0;
}