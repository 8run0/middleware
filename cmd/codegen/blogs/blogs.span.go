package blogs

		var _ bloggerImpl = &bloggerSpanner{}

		type bloggerSpanner struct {
			*OTELTools
			next bloggerImpl
		}
		 
	    
		func (s *bloggerSpanner) addBlog ( req AddBlogRequest,) (res AddBlogResponse,) {
			ctx, span := s.Tracer.Start(s.Ctx, "blogger_addBlog")
			s.Ctx = ctx
			defer span.End()
			return s.next.addBlog( req,) 
		}
		 
	    
		func (s *bloggerSpanner) deleteBlog ( req DeleteBlogRequest,) (res DeleteBlogResponse,) {
			ctx, span := s.Tracer.Start(s.Ctx, "blogger_deleteBlog")
			s.Ctx = ctx
			defer span.End()
			return s.next.deleteBlog( req,) 
		}
		 
	    
		func (s *bloggerSpanner) listBlog ( req ListBlogRequest,) (res ListBlogResponse,) {
			ctx, span := s.Tracer.Start(s.Ctx, "blogger_listBlog")
			s.Ctx = ctx
			defer span.End()
			return s.next.listBlog( req,) 
		}
		 
	    
		func (s *bloggerSpanner) getBlog ( req GetBlogRequest,) (res GetBlogResponse,) {
			ctx, span := s.Tracer.Start(s.Ctx, "blogger_getBlog")
			s.Ctx = ctx
			defer span.End()
			return s.next.getBlog( req,) 
		}
		

		