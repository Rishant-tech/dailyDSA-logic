#include <bits/stdc++.h>
using namespace std;

bool checkInclusion(string s1,string s2){
     vector<int> freq1(26,0);
     vector<int> freq2(26,0);
        for(char ch:s1){
            freq1[ch-'a']++;
        }
        int i=0,j=0;
        while(j<s2.length()){
            freq2[s2[j]-'a']++;
            if(j-i+1 == s1.length()){
                if(freq1==freq2){
                    return true;
                }else{
                    freq2[s2[i]-'a']--;
                    i++;
                }
            }
            j++;
        }
        return false;
}

int main(){
    string s1,s2;
    cout<<"enter string 1";
    cin>>s1;
    cout<<"enter string 2";
    cin>>s2;
    bool ans = checkInclusion(s1,s2);
    cout<<ans;
}