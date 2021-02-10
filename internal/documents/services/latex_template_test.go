package services

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/models"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestLatexTemplate_Inject_SimpleTlp(t *testing.T) {
	testDate, err := time.Parse("2006-01-02", "2020-09-27")
	if err != nil {
		panic(err)
	}
	testHeader := models.MemoForRecordHeader{
		Tlp:           models.GREEN,
		ControlNumber: "MR-00-000",
		Date:          &models.Date{Time: testDate},
		MemoFor:       "RECORD---STANDARDS TRACK",
		MemoFrom:      "Robert Hawk",
		Subject:       "Test Memorandum",
		Attachments:   []string{"Thing 1", "Thing 2"},
		Cc:            []string{"Robert Hawk"},
	}

	testText := `\hypertarget{section}{%
\paragraph{}\label{section}}

This is a test memorandum.

\hypertarget{section-1}{%
\paragraph{}\label{section-1}}

For use in testing things.

\hypertarget{section-2}{%
\paragraph{}\label{section-2}}

By Testificates.
`

	testTemplate := `\documentclass[
    %TLP%
]{thusmemo}

\usepackage{amssymb}
\usepackage{arevmath}
\usepackage{booktabs}
\usepackage[hidelinks]{hyperref}
\usepackage{longtable}
\usepackage{pifont}

\providecommand{\tightlist}{\setlength{\itemsep}{0pt}\setlength{\parskip}{0pt}}

\newcommand{\cmark}{\ding{51}}
\newcommand{\Unchecked}{$\square$}
\newcommand{\Checked}{\rlap{$\square$}{\large\hspace{1pt}\cmark}}
\newcommand{\Blank}{\rule{5em}{1pt}}

\date{%
    %CONTROL_NUMBER%
    \\
    %DATE%
}

\memofor{%
    %MEMO_FOR%
}

\memofrom{%
    %MEMO_FROM%
}

\subject{%
    %SUBJECT%
}


\begin{document}
    \maketitle

    %TEXT%

    \digisign{Robert H. Hawk}{}{}

    %ATTACHMENTS%

    %CC%

    %STATUS%

\end{document}
`

	expected := `\documentclass[
    green
]{thusmemo}

\usepackage{amssymb}
\usepackage{arevmath}
\usepackage{booktabs}
\usepackage[hidelinks]{hyperref}
\usepackage{longtable}
\usepackage{pifont}

\providecommand{\tightlist}{\setlength{\itemsep}{0pt}\setlength{\parskip}{0pt}}

\newcommand{\cmark}{\ding{51}}
\newcommand{\Unchecked}{$\square$}
\newcommand{\Checked}{\rlap{$\square$}{\large\hspace{1pt}\cmark}}
\newcommand{\Blank}{\rule{5em}{1pt}}

\date{%
    MR--00--000
    \\
    27 September 2020
}

\memofor{%
    RECORD---STANDARDS TRACK
}

\memofrom{%
    Robert Hawk
}

\subject{%
    Test Memorandum
}


\begin{document}
    \maketitle

    \hypertarget{section}{%
\paragraph{}\label{section}}

This is a test memorandum.

\hypertarget{section-1}{%
\paragraph{}\label{section-1}}

For use in testing things.

\hypertarget{section-2}{%
\paragraph{}\label{section-2}}

By Testificates.

    \digisign{Robert H. Hawk}{}{}

    \attachments{%
\item Thing 1
\item Thing 2
}

    \cc{%
\item Robert Hawk
}

    %STATUS%

\end{document}
`

	templ := NewLatexTemplate([]byte(testTemplate))
	result, err := templ.Inject(testHeader, []byte(testText))

	assert.NoError(t, err)
	assert.Equal(t, strings.Join(strings.Fields(expected), ""), strings.Join(strings.Fields(string(result)), ""))
}

/*
func TestLatexTemplate_Inject_TlpWithCaveats(t *testing.T) {
	testDate, err := time.Parse("2006-01-02", "2020-09-27")
	if err != nil {
		panic(err)
	}
	testHeader := models.MemoForRecordHeader{
		Tlp:           models.RED,
		Caveats:       []string{"FAMILY", "FRIENDS"},
		ControlNumber: "MR-00-000",
		Date:          &models.Date{Time: testDate},
		MemoFor:       "RECORD---STANDARDS TRACK",
		MemoFrom:      "Robert Hawk",
		Subject:       "Test Memorandum",
		Attachments:   []string{"Thing 1", "Thing 2"},
		Cc:            []string{"Robert Hawk"},
	}

	testContent := `\hypertarget{section}{%
\paragraph{}\label{section}}

This is a test memorandum.

\hypertarget{section-1}{%
\paragraph{}\label{section-1}}

For use in testing things.

\hypertarget{section-2}{%
\paragraph{}\label{section-2}}

By Testificates.
`

	testTemplate := `\documentclass[
    %TLP%
]{thusmemo}

\usepackage{amssymb}
\usepackage{arevmath}
\usepackage{booktabs}
\usepackage[hidelinks]{hyperref}
\usepackage{longtable}
\usepackage{pifont}

\providecommand{\tightlist}{\setlength{\itemsep}{0pt}\setlength{\parskip}{0pt}}

\newcommand{\cmark}{\ding{51}}
\newcommand{\Unchecked}{$\square$}
\newcommand{\Checked}{\rlap{$\square$}{\large\hspace{1pt}\cmark}}
\newcommand{\Blank}{\rule{5em}{1pt}}

\date{%
    %CONTROL_NUMBER%
    \\
    %DATE%
}

\memofor{%
    %MEMO_FOR%
}

\memofrom{%
    %MEMO_FROM%
}

\subject{%
    %SUBJECT%
}


\begin{document}
    \maketitle

    %TEXT%

    \digisign{Robert H. Hawk}{}{}

    %ATTACHMENTS%

    %CC%

    %STATUS%

\end{document}
`
	expected := `\documentclass[
    %TLP%
]{thusmemo}

\usepackage{amssymb}
\usepackage{arevmath}
\usepackage{booktabs}
\usepackage[hidelinks]{hyperref}
\usepackage{longtable}
\usepackage{pifont}

\providecommand{\tightlist}{\setlength{\itemsep}{0pt}\setlength{\parskip}{0pt}}

\newcommand{\cmark}{\ding{51}}
\newcommand{\Unchecked}{$\square$}
\newcommand{\Checked}{\rlap{$\square$}{\large\hspace{1pt}\cmark}}
\newcommand{\Blank}{\rule{5em}{1pt}}

\date{%
    %CONTROL_NUMBER%
    \\
    %DATE%
}

\memofor{%
    %MEMO_FOR%
}

\memofrom{%
    %MEMO_FROM%
}

\subject{%
    %SUBJECT%
}


\begin{document}
    \maketitle

    %TEXT%

    \digisign{Robert H. Hawk}{}{}

    %ATTACHMENTS%

    %CC%

    %STATUS%

\end{document}
`
}


 */