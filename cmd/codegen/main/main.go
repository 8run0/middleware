package main

	type Calculator struct {
		calculatorImpl
		*OTELTools
	}

	func NewCalculator(tools *OTELTools) *Calculator {
		
		calculator := calculatorSpanner{
			OTELTools: tools,
			next:  &calculator{},
		}
		return &Calculator{
			calculatorImpl: &calculator,
		}
	}

	var _ calculatorImpl = &calculator{}

	type calculator struct {
	}
	 
	    
		func (*calculator) add ( x int64, y int64,) (z int64,) {
			//add business logic goes here
			return
		}
		 
	    
		func (*calculator) sub ( x int64, y int64,) (z int64,) {
			//sub business logic goes here
			return
		}
		 
	    
		func (*calculator) multiply ( x int64, y int64,) (z int64,) {
			//multiply business logic goes here
			return
		}
		 
	    
		func (*calculator) divide ( x int64, y int64,) (z int64,) {
			//divide business logic goes here
			return
		}
		
	