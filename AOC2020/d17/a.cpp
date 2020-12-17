#include <iostream>
#include <vector>

using namespace std;
 
#define fr(i, st, n) for(int i = (int )st; i < (int )n; i++)
#define rfr(i, en, st) for(int i = (int )en; i >= (int )st; i--)
#define pb push_back
#define sz(a) (int)(a).size()
 
typedef vector< int> vi;

 // START FROM HERE

const int lim=30;
const int sx=11,sy=11,sz=11;
bool grid[lim][lim][lim];

struct Coo
{
 int x,y,z;
 // Coo(){}
 Coo(int x,int y,int z):x(x),y(y),z(z){}
};

int cycle(){
 std::vector<Coo> v;
 auto f=[](int x){
	return x>=0&&x<lim;
 };
 int tot=0;
 fr(x,0,lim){
	fr(y,0,lim){
	 fr(z,0,lim){
		int cnt=0;
		fr(dx,-1,2)fr(dy,-1,2)fr(dz,-1,2){
		 if(dx==0&&dy==0&&dz==0)continue;
		 if(f(dx+x)&&f(dy+y)&&f(dz+z)&&grid[x+dx][y+dy][z+dz]){
			++cnt;
		 }
		}
		if(grid[x][y][z]){
		 tot++;
		 if (cnt==2||cnt==3){

		 }else{
			v.pb(Coo(x,y,z));
		 }
		}
		if(!grid[x][y][z]&&cnt==3){
		 v.pb(Coo(x,y,z));
		}

	 }
	}
 }
 for(auto c:v){
	grid[c.x][c.y][c.z] = !grid[c.x][c.y][c.z];
	if (grid[c.x][c.y][c.z]) tot++;
	else tot--;
 }
 return tot;


}

int main(int argc, char const *argv[])
{
 fr(i,0,lim)fr(j,0,lim)fr(k,0,lim){
	grid[i][j][k]=false;
 }
 // std::vector<string> mat({".#.","..#","###"});
 std::vector<string> mat({"..#....#",
"##.#..##",
".###....",
"#....#.#",
"#.######",
"##.#....",
"#.......",
".#......"});
 fr(i,0,sz(mat)){
	fr(j,0,sz(mat[0])){
	 if (mat[i][j]=='#'){
	 grid[i+sx][j+sy][sz]=true;
	 }
	}
 }
 fr(i,0,6)
 cout<<cycle()<<endl;


 return 0;
}





// {"..#....#",
// "##.#..##",
// ".###....",
// "#....#.#",
// "#.######",
// "##.#....",
// "#.......",
// ".#......"}