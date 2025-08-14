package main

import (
    "errors"
    "fmt"
    "strconv"
    "strings"
)

const (
    RESULT_EQ = 1 // Equal
    RESULT_LT = 2 // Less Than
    RESULT_GT = 3 // Greater Than

    TYPE_NUM    = 1
    TYPE_LETTER = 2
    TYPE_SPOT   = 3
    TYPE_OTHER  = 4
)

// GetType 判断字符类型
func GetType(v rune) uint8 {
    switch {
    case '0' <= v && v <= '9':
        return TYPE_NUM
    case ('a' <= v && v <= 'z') || ('A' <= v && v <= 'Z'):
        return TYPE_LETTER
    case v == '.':
        return TYPE_SPOT
    default:
        return TYPE_OTHER
    }
}

// paramsCheck 参数校验：检查版本号是否合法
func paramsCheck(str string) error {
    if str == "" {
        return errors.New("invalid params")
    }
    if str[0] == '.' {
        return fmt.Errorf("版本号不能以 . 开头: %s", str)
    }
    if str[len(str)-1] == '.' {
        return fmt.Errorf("版本号不能以 . 结尾: %s", str)
    }
    for i := range str {
        if GetType(rune(str[i])) == TYPE_OTHER {
            return fmt.Errorf("版本号包含非法字符: %s", str)
        }
    }
    return nil
}

// compareSegment 比较两个版本段
func compareSegment(a, b string) int8 {
    // 尝试转为数字比较
    na, errA := strconv.Atoi(a)
    nb, errB := strconv.Atoi(b)

    if errA == nil && errB == nil {
        if na > nb {
            return RESULT_GT
        } else if na < nb {
            return RESULT_LT
        }
        return RESULT_EQ
    }

    // 否则按字符串比较
    if a > b {
        return RESULT_GT
    } else if a < b {
        return RESULT_LT
    }
    return RESULT_EQ
}

// VersionCompare 比较两个版本号字符串
func VersionCompare(one, two string) (int8, error) {
    if err := paramsCheck(one); err != nil {
        return 0, err
    }
    if err := paramsCheck(two); err != nil {
        return 0, err
    }
    if one == two {
        return RESULT_EQ, nil
    }

    // 按 . 分割成段
    oneSegs := strings.Split(one, ".")
    twoSegs := strings.Split(two, ".")

    // 逐段比较
    for i := 0; i < len(oneSegs) && i < len(twoSegs); i++ {
        res := compareSegment(oneSegs[i], twoSegs[i])
        if res != RESULT_EQ {
            return res, nil
        }
    }

    // 所有段都相同，长度不同则更长者更大
    if len(oneSegs) > len(twoSegs) {
        return RESULT_GT, nil
    } else if len(oneSegs) < len(twoSegs) {
        return RESULT_LT, nil
    }

    return RESULT_EQ, nil
}