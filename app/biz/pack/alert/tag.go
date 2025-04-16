package alert

import (
	"darkhawk/app/biz/model"
	"darkhawk/pkg/utils"
	"regexp"
	"strings"
)

// TagRewrite 标签重写
func TagRewrite(tags map[string]string, tagRewrites []model.TagRewrite) map[string]string {
	if tagRewrites == nil {
		return tags
	}
	for _, v := range tagRewrites {
		// 标签不存在，不进行操作
		if _, ok := tags[v.RewriteItem.OldTag]; !ok {
			continue
		}
		// 没有匹配过滤条件时，不进行操作
		if v.Filters != nil && !HasTag(v.Filters, tags) {
			continue
		}

		switch v.RewriteType {
		case model.TagTypeExtract:
			tags = ExtractTag(tags, v)
		case model.TagTypeRewrite:
			tags = RewriteTag(tags, v)
		case model.TagTypeDelete:
			tags = DeleteTag(tags, v.RewriteItem.DeleteTag)
		}
	}
	return tags
}

// ExtractTag 提取标签
func ExtractTag(tags map[string]string, definition model.TagRewrite) map[string]string {
	re := regexp.MustCompile(definition.RewriteItem.Expression)
	matches := re.FindStringSubmatch(tags[definition.RewriteItem.OldTag])
	if len(matches) == 1 {
		tags[definition.RewriteItem.NewTag] = matches[0]
		return tags
	} else if len(matches) >= 2 {
		tags[definition.RewriteItem.NewTag] = matches[1]
		return tags
	}
	return tags
}

// RewriteTag 标签重写
func RewriteTag(tags map[string]string, definition model.TagRewrite) map[string]string {
	s := tags[definition.RewriteItem.OldTag]
	re := regexp.MustCompile(definition.RewriteItem.Expression)
	matches := re.FindStringSubmatch(tags[definition.RewriteItem.OldTag])

	if len(matches) == 1 {
		tags[definition.RewriteItem.NewTag] = strings.ReplaceAll(s, definition.RewriteItem.Expression, definition.RewriteItem.Value)
		return tags
	} else if len(matches) >= 2 {
		tags[definition.RewriteItem.NewTag] = re.ReplaceAllString(s, definition.RewriteItem.Value)
		return tags
	}
	return tags
}

// DeleteTag 标签删除
func DeleteTag(tags map[string]string, deleteTag []string) map[string]string {
	newTags := make(map[string]string)
	for k, v := range tags {
		if !utils.InOfT(k, deleteTag) {
			newTags[k] = v
		}
	}
	return newTags
}

// HasTag 标签是否匹配
func HasTag(filter []model.TagFilter, tags map[string]string) bool {
	for _, f := range filter {
		pass := true
		if f.Operation == "IN" {
			if _, ok := tags[f.Tag]; !ok {
				pass = false
			}
			in := false
			for _, item := range f.Value {
				if strings.Contains(tags[f.Tag], item) {
					in = true
				}
			}
			if !in {
				pass = false
			}

		} else if f.Operation != "IN" {
			if _, ok := tags[f.Tag]; !ok {
				pass = true
			}
			in := false
			for _, item := range f.Value {
				if strings.Contains(tags[f.Tag], item) {
					in = true
				}
			}
			if in {
				pass = false
			}
		}
		if !pass {
			return false
		}
	}
	return true
}
