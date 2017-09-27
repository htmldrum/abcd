package divisions

type Division struct {
	Name string
}

var (
	News  = Division{"News"}
	Tv    = Division{"Tv"}
	Radio = Division{"Radio"}
	List  = []Division{News, Tv, Radio}
)
