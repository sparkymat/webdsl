webdsl
======
DSL for generating HTML (TBD) and CSS.

**Sample usage for CSS**

	  Css("main").Rules(
	  	CssRules().For(".modal").Set(
	  		FontSizeEm(4),
	  		Color("#ffffff"),
	  		BackgroundColor("#000000"),
	  	),
	  	CssRules().For("#create-button").For("#update-button").Set(
	  		BorderRadius(6),
	  		FontFamily("PT Sans", "serif"),
	  		BackgroundImage("/images/button.png"),
	  		BackgroundRepeat(CssBackgroundRepeatX),
	  	),
	  	CssRules().For(".margin-rules").Set(
	  		Margin1a(4),
	  		Margin2a(6, 6),
	  		Margin3a(8, 8, 8),
	  		Margin4a(10, 10, 10, 10),
	  		MarginLeft(2),
	  		MarginRight(2),
	  		MarginTop(2),
	  		MarginBottom(2),
	  	),
	  )

