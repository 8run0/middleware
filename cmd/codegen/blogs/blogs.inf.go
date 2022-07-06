package blogs

	type bloggerImpl interface {
		addBlog ( req AddBlogRequest,) (res AddBlogResponse,)
		deleteBlog ( req DeleteBlogRequest,) (res DeleteBlogResponse,)
		listBlog ( req ListBlogRequest,) (res ListBlogResponse,)
		getBlog ( req GetBlogRequest,) (res GetBlogResponse,)
		
	}
	