package funyu

import (
	"testing"
)

func test(t *testing.T, err error) {
	if err != nil {
		t.Error(err.Error())
	}
}

func TestBlock(t *testing.T) {
	s := NewSection("test")
	test(t, s.Feed("hello"))
	test(t, s.Feed(""))
	test(t, s.Feed("-- hoge"))
	test(t, s.Feed("\tfuga"))
	test(t, s.Feed(""))
	test(t, s.Feed("-- empty"))
	test(t, s.Feed(""))
	test(t, s.Feed("p.s. 2016-01-01"))
	test(t, s.Feed("\tpost script"))
	test(t, s.Feed(""))
	test(t, s.Feed("\ttest"))
	test(t, s.Feed("\tabc"))
	test(t, s.Feed(""))
	test(t, s.Feed("``` HTML"))
	test(t, s.Feed("\tthis is <b>html</b>"))
	test(t, s.Feed("```"))
	test(t, s.Feed(""))
	test(t, s.Feed("((("))
	test(t, s.Feed("\tembedded <em>HTML</em>"))
	test(t, s.Feed(")))"))
	test(t, s.Feed(""))

	h := s.HTML(1)
	if h != `<section>
<h1>test</h1>
<p>
hello<br>
</p>
<section>
<h2>hoge</h2>
<p>
fuga<br>
</p>
</section>
<section>
<h2>empty</h2>
</section>
<ins>
<b>p.s. <date>2016-01-01</date></b>
<p>
post script<br>
</p>
<p>
test<br>
abc<br>
</p>
</ins>
<pre class="code" data-language=HTML>this is &lt;b&gt;html&lt;/b&gt;</pre>
embedded <em>HTML</em>
</section>
` {
		t.Error("Failed generate HTML:\n"+h)
	}

	h = s.HTML(3)
	if h != `<section>
<h3>test</h3>
<p>
hello<br>
</p>
<section>
<h4>hoge</h4>
<p>
fuga<br>
</p>
</section>
<section>
<h4>empty</h4>
</section>
<ins>
<b>p.s. <date>2016-01-01</date></b>
<p>
post script<br>
</p>
<p>
test<br>
abc<br>
</p>
</ins>
<pre class="code" data-language=HTML>this is &lt;b&gt;html&lt;/b&gt;</pre>
embedded <em>HTML</em>
</section>
` {
		t.Error("Failed change level:\n"+h)
	}
}
