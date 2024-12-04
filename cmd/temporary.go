package cmd

type Command struct{	
	Use string
	Short string
	Run func(args []string)
}

func main(){

}