package thughash

import(
	"testing"
)

type testcase struct{
	test float64
	answer string 
}

type degenTest struct{
	test string
	answer float64
}

type genFrom struct{
	test string
	answer string
}

type thugTest struct{
	test int64
	subject string
	verb string
	end string
	remainder int
	quickhash string
	answer string 
}

type testInts struct{
	test int64
	answer int
}

func TestThirtyTwoEncode(t *testing.T) {
    t.Log("Testing the ThirtyTwoEncode method")
    var th ThugHash
    testArray := []testcase{
    	testcase{15, "000000f"},
    	testcase{999, "00000v7"},
    	testcase{9999, "00009of"},
    	testcase{0, "0000000"},
    	testcase{2147483647, "1vvvvvv"},
    	testcase{2147483648, "2000000"}}
    for i := 0; i < len(testArray); i++ {
	    x := th.ThirtyTwoEncode(testArray[i].test)
	    if x != testArray[i].answer{
	        t.Errorf("%d doesn't equal %q.", x, testArray[i].answer)
	    }
	}
}

func TestGenerate(t *testing.T) {
	t.Log("Testing the Generate method")
	testArray := []testcase{
		testcase{0,"Baby-actin-a-double-cup-1000"},
		testcase{227425999,"Mothafucka-turnt-holdin-it-down-1867"},
    	testcase{15, "Baby-actin-holdin-it-down-1000"},
    	testcase{999, "Baby-escapin-squad-1000"},
    	testcase{9999, "Balla-pimpin-holdin-it-down-1000"},
    	testcase{2147483647, "Weezy-woke-up-yo-ass-9191"},
    	testcase{2147483648, "Baby-actin-a-double-cup-9192"},
		testcase{247333,"Turnt-interceptin-skrilla-1000"}}
	for i := 0; i < len(testArray); i++ {
		var th ThugHash	
		th.Generate(testArray[i].test)
	    x := th.MakeSlug()
	    if x != testArray[i].answer{
	        t.Errorf("%q doesn't equal %q.", x, testArray[i].answer)
	    }
	}
}

func TestGenerateRemainder(t *testing.T) {
	t.Log("Testing the GenerateRemainder method")
	var th ThugHash
	testArray := []testInts{
		testInts{1,1001},
		testInts{9999,9999},
		testInts{8999,9999},
		testInts{404,1404},
		testInts{0,1000}}
    for i := 0; i < len(testArray); i++ {
	    x := th.GenerateRemainder(testArray[i].test)
	    if x != testArray[i].answer{
	        t.Errorf("%d doesn't equal %d.", x, testArray[i].answer)
	    }
	}	
}

func TestMakeSlug(t *testing.T) {
	t.Log("Testing the MakeSlug method")
	var th ThugHash
	testArray := []thugTest{
		thugTest{1,"OG","sippin","my-stash",1449,"sdfsdf","OG-sippin-my-stash-1449"},
		thugTest{2,"Yayo","poppin","the-building",1926,"sdfsdf","Yayo-poppin-the-building-1926"},
		thugTest{3,"Yayo","poppin","the-building",2005,"sdfsdf","Yayo-poppin-the-building-2005"}}
    for i := 0; i < len(testArray); i++ {
    	th.Index = int(testArray[i].test)
    	th.Subject = testArray[i].subject
    	th.Verb = testArray[i].verb
    	th.End = testArray[i].end
    	th.Remainder = testArray[i].remainder
    	th.QuickHash = testArray[i].quickhash
	    x := th.MakeSlug()
	    if x != testArray[i].answer{
	        t.Errorf("%q doesn't equal %q.", x, testArray[i].answer)
	    }
	}	
}

func TestMatchHash(t *testing.T) {
	t.Log("Testing the MatchHash method")
	var th ThugHash
	testArray := []thugTest{
		thugTest{1,"Balla","pimpin","holdin-it-down",1000,"00009of","Balla-pimpin-holdin-it-down-1000"},
		thugTest{2,"Weezy","woke-up","yo-ass",9191,"1vvvvvv","Weezy-woke-up-yo-ass-9191"},
		thugTest{3,"Baby","actin","a-double-cup",9192,"2000000","Baby-actin-a-double-cup-9192"}}
	for i := 0; i < len(testArray); i++ {
		th.Index = int(testArray[i].test)
    	th.Subject = testArray[i].subject
    	th.Verb = testArray[i].verb
    	th.End = testArray[i].end
    	th.Remainder = testArray[i].remainder
    	th.QuickHash = testArray[i].quickhash
	    x := th.MatchHash(th.MakeSlug())
	    if x != testArray[i].quickhash{
	        t.Errorf("%q doesn't equal %q.", x, testArray[i].quickhash)
	    }
	    y := th.MatchHash(th.QuickHash)
	    if y != testArray[i].answer{
	        t.Errorf("%q doesn't equal %q.", y, testArray[i].answer)
	    }
	}
}

func TestDegenerate(t *testing.T) {
	t.Log("Testing the Degenerate method")
	var th ThugHash
	testArray := []degenTest{
		degenTest{"Balla-pimpin-holdin-it-down-1000",9999},
		degenTest{"Weezy-woke-up-yo-ass-9191",2147483647},
		degenTest{"Yayo-poppin-the-building-2005",2.63718831e+08}}
	for i := 0; i < len(testArray); i++ {
		x := th.Degenerate(testArray[i].test)
		if x != testArray[i].answer{
	        t.Errorf("%v doesn't equal %v.", x, testArray[i].answer)
	    }
	}	
}

func TestGenerateFrom(t *testing.T){
	t.Log("Testing the GenerateFrom method")
	var th ThugHash
	testArray := []genFrom{
		genFrom{"1vvvvvv","1vvvvvv"},
		genFrom{"0000000","0000000"},
		genFrom{"Baby-actin-a-double-cup-9192","2000000"}}
	for i := 0; i < len(testArray); i++ {
		th.GenerateFrom(testArray[i].test)
		x := th.QuickHash
		if x != testArray[i].answer{
	        t.Errorf("%v doesn't equal %v.", x, testArray[i].answer)
	    }
	}
}