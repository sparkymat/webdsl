webdsl
======
DSL for generating HTML (TBD) and CSS.

**Sample usage for CSS**

	  Css("main").Rules(
	  	Rule().For(".modal").Set(
	  		FontSizeEm(4),
	  		Color("#ffffff").Important(),
	  		BackgroundColor("#000000"),
	  	),
	  	Rule().For("#create-button").For("#update-button").Set(
	  		BorderRadius(6),
	  		FontFamily("PT Sans", "serif"),
	  		BackgroundImage("/images/button.png"),
	  		BackgroundRepeat(CssBackgroundRepeatX),
	  	),
	  	Rule().For(".margin-rules").Set(
	  		Margin1i(4),
	  		Margin2i(6, 6),
	  		Margin3i(8, 8, 8),
	  		Margin4i(10, 10, 10, 10),
	  		MarginLeft(2),
	  		MarginRight(2),
	  		MarginTop(2),
	  		MarginBottom(2),
	  	),
	  )

