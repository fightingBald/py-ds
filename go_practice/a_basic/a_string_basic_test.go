package a_basic

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func TestStringOps(t *testing.T) {
	t.Run("CRUD_like_ops", func(t *testing.T) {
		// Create

		s := "  Hello, 世界  "
		// Read: 长度、是否包含、位置
		lens := len(s)
		containHello := strings.Contains(s, "hello")
		helloIndex := strings.Index(s, "hello")
		pythonIndex := strings.Index(s, "python")

		require.Equal(t, 2+7+2, lens) // 字节长度（注意中文是多字节）
		require.True(t, containHello)
		require.Equal(t, 2, helloIndex) // 第一次出现位置
		require.Equal(t, -1, pythonIndex)

		// Update（都是返回新串）
		s1 := strings.TrimSpace(s)
		s2 := strings.ReplaceAll(s1, "Hello", "Hi")
		s3 := strings.ToUpper(s2)
		s4 := strings.ToLower(s3)
		require.Equal(t, "hi, 世界", s4)

		// Delete（本质也是构造新串）：删除逗号
		s5 := strings.ReplaceAll(s4, ",", "")
		require.Equal(t, "hi 世界", s5)

		// Substring（切片，按字节，UTF-8 小心）：这里截取前 2 字节 "hi"
		firstTwo := s5[:2]
		require.Equal(t, "hi", firstTwo)

		// 遍历：按 rune（正确处理中文）
		var runes []rune
		for _, r := range s5 {
			runes = append(runes, r)
		}
		require.True(t, utf8.ValidString(s5))
		require.GreaterOrEqual(t, len(runes), 3)

		// Split / Join
		parts := strings.Split("go,python,java", ",")
		joined := strings.Join(parts, "|")

		require.Equal(t, []string{"go", "python", "java"}, parts)
		require.Equal(t, "go|python|java", joined)

		//格式化
		name := "Huayi"
		age := 26
		formatted := fmt.Sprintf("Name: %s, Age: %d", name, age)
		fmt.Println(formatted)

		// 类 map/filter：先 split 得切片，再函数式处理
		langs := strings.Split("go,python,java,js", ",")
		upper := MapSlice(langs, func(s string) string { return strings.ToUpper(s) })
		require.Equal(t, []string{"GO", "PYTHON", "JAVA", "JS"}, upper)

		filtered := FilterSlice(langs, func(s string) bool { return len(s) <= 2 })
		require.Equal(t, []string{"go", "js"}, filtered)
	})

	t.Run("real_world_text_cleaning", func(t *testing.T) {
		raw := "  product:  iPhone 15 Pro  , price:  999  "
		clean := strings.TrimSpace(raw)
		clean = strings.ReplaceAll(clean, "  ", " ")
		clean = strings.ReplaceAll(clean, " ,", ",")
		require.Equal(t, "product: iPhone 15 Pro, price: 999", clean)
	})
}
