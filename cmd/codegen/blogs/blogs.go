package blogs

	type Blogger struct {
		bloggerImpl
		*OTELTools
	}

	func NewBlogger(tools *OTELTools) *Blogger {
		
		blogger := bloggerSpanner{
			OTELTools: tools,
			next:  &blogger{},
		}
		return &Blogger{
			bloggerImpl: &blogger,
		}
	}

	var _ bloggerImpl = &blogger{}

	type blogger struct {
	}
	 
	    
		func (*blogger) addBlog ( req AddBlogRequest,) (res AddBlogResponse,) {
			//addBlog business logic goes here
			return
		}
		 
	    
		func (*blogger) deleteBlog ( req DeleteBlogRequest,) (res DeleteBlogResponse,) {
			//deleteBlog business logic goes here
			return
		}
		 
	    
		func (*blogger) listBlog ( req ListBlogRequest,) (res ListBlogResponse,) {
			//listBlog business logic goes here
			return
		}
		 
	    
		func (*blogger) getBlog ( req GetBlogRequest,) (res GetBlogResponse,) {
			//getBlog business logic goes here
			return
		}
		
	