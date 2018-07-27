package data

import (
	"regexp"
	"strings"
)

type TTagsProcessor struct{}

var TagsProcessor *TTagsProcessor = &TTagsProcessor{}
var normalizeTagRegex = regexp.MustCompile("(_|#)+")
var compressTagRegex = regexp.MustCompile("( |_|#)")
var splitTagRegex = regexp.MustCompile("(,|;)+")
var hashTagRegex = regexp.MustCompile("#\\w+")

func (c *TTagsProcessor) NormalizeTag(tag string) string {
	tag = normalizeTagRegex.ReplaceAllString(tag, " ")
	return strings.Trim(tag, " \t\r\n")
}

func (c *TTagsProcessor) CompressTag(tag string) string {
	tag = compressTagRegex.ReplaceAllString(tag, "")
	return strings.ToLower(tag)
}

func (c *TTagsProcessor) EqualTags(tag1 string, tag2 string) bool {
	if tag1 == "" && tag2 == "" {
		return true
	}
	if tag1 == "" || tag2 == "" {
		return false
	}
	return strings.Compare(c.CompressTag(tag1), c.CompressTag(tag2)) == 0
}

func (c *TTagsProcessor) NormalizeTags(tags []string) []string {
	for index := 0; index < len(tags); index++ {
		tags[index] = c.NormalizeTag(tags[index])
	}
	return tags
}

func (c *TTagsProcessor) NormalizeTagList(tagList string) []string {
	tags := splitTagRegex.Split(tagList, -1)
	return c.NormalizeTags(tags)
}

func (c *TTagsProcessor) CompressTags(tags []string) []string {
	for index := 0; index < len(tags); index++ {
		tags[index] = c.CompressTag(tags[index])
	}
	return tags
}

func (c *TTagsProcessor) CompressTagList(tagList string) []string {
	tags := splitTagRegex.Split(tagList, -1)
	return c.CompressTags(tags)
}

func (c *TTagsProcessor) ExtractHashTags(text string) []string {
	if text == "" {
		return []string{}
	}

	hashTags := hashTagRegex.FindAllString(text, -1)
	tags := make([]string, 0, len(hashTags))

	// Filter duplicates
	for index := 0; index < len(hashTags); index++ {
		tag := c.CompressTag(hashTags[index])
		duplicate := false
		for index2 := 0; index2 < len(tags); index2++ {
			if strings.Compare(tags[index2], tag) == 0 {
				duplicate = false
				break
			}
		}
		if !duplicate {
			tags = append(tags, tag)
		}
	}

	return tags
}

// private static extractString(field: any): string {
// 	if (field == null) return '';
// 	if (_.isString(field)) return field;
// 	if (!_.isObject(field)) return '';

// 	let result = '';
// 	for (let prop in field) {
// 		result += ' ' + TagsProcessor.extractString(field[prop]);
// 	}
// 	return result;
// }

// public static extractHashTags(obj: any, ...searchFields: string[]): string[] {
// 	let tags = TagsProcessor.compressTags(obj.tags);

// 	_.each(searchFields, (field) => {
// 		let text = TagsProcessor.extractString(obj[field]);

// 		if (text != '') {
// 			let hashTags = text.match(TagsProcessor.HASHTAG_REGEX);
// 			tags = tags.concat(TagsProcessor.compressTags(hashTags));
// 		}
// 	});

// 	return _.uniq(tags);
// }
