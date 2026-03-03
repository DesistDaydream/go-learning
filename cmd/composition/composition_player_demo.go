package main

import "fmt"

// ============================================
// 痛点1：继承导致的紧耦合问题
// ============================================

// 假设我们有一个游戏，有各种角色和装备

// 使用组合前：如果用继承，可能需要这样的结构
// Character -> Warrior -> ArmedWarrior
//           -> Mage -> ArmedMage
// 这样会导致类爆炸，而且武器逻辑重复

// ============================================
// 解决方案：使用组合
// ============================================

// Weapon 武器接口
type Weapon interface {
	Attack() string
	Damage() int
}

// Sword 剑
type Sword struct {
	Name  string
	Power int
}

func (s Sword) Attack() string {
	return fmt.Sprintf("用%s挥砍", s.Name)
}

func (s Sword) Damage() int {
	return s.Power
}

// Bow 弓
type Bow struct {
	Name  string
	Power int
}

func (b Bow) Attack() string {
	return fmt.Sprintf("用%s射击", b.Name)
}

func (b Bow) Damage() int {
	return b.Power
}

// Armor 防具接口
type Armor interface {
	Defend() string
	Defense() int
}

// SteelArmor 钢甲
type SteelArmor struct {
	Name       string
	Protection int
}

func (s SteelArmor) Defend() string {
	return fmt.Sprintf("用%s格挡攻击", s.Name)
}

func (s SteelArmor) Defense() int {
	return s.Protection
}

// ============================================
// 核心：Character 通过组合拥有武器和防具能力
// ============================================

// Character 角色 - 通过组合获得功能
type Character struct {
	Name   string
	Health int
	// 组合：角色"拥有"武器和防具，而不是"是"武器或防具
	Weapon Weapon // 嵌入接口类型
	Armor  Armor  // 嵌入接口类型
}

// Fight 战斗方法
func (c *Character) Fight(target string) {
	if c.Weapon != nil {
		fmt.Printf("%s %s，造成%d点伤害！\n", c.Name, c.Weapon.Attack(), c.Weapon.Damage())
	} else {
		fmt.Printf("%s 没有武器，只能徒手攻击！\n", c.Name)
	}
}

// Defend 防御方法
func (c *Character) Defend() {
	if c.Armor != nil {
		fmt.Printf("%s %s，减少%d点伤害\n", c.Name, c.Armor.Defend(), c.Armor.Defense())
	} else {
		fmt.Printf("%s 没有防具，只能硬抗！\n", c.Name)
	}
}

// ============================================
// 痛点2：需要灵活组合多个功能
// ============================================

// 使用匿名嵌入（Embedding）实现更简洁的组合
type Hero struct {
	Name string
	// 匿名嵌入：直接继承类型的方法
	Sword      // 匿名嵌入 Sword 类型
	SteelArmor // 匿名嵌入 SteelArmor 类型
}

// Hero 自动获得了 Sword 和 SteelArmor 的方法

// ============================================
// 痛点3：运行时动态更换能力
// ============================================

// 由于使用组合，我们可以在运行时更换装备
func (c *Character) EquipWeapon(w Weapon) {
	c.Weapon = w
	fmt.Printf("%s 装备了 %T\n", c.Name, w)
}

func (c *Character) EquipArmor(a Armor) {
	c.Armor = a
	fmt.Printf("%s 装备了 %T\n", c.Name, a)
}

func main() {
	fmt.Println("=== Go 语言组合（Composition）实践 ===")

	// 示例1：基础组合
	fmt.Println("【示例1：基础组合】")
	warrior := &Character{
		Name:   "战士",
		Health: 100,
		Weapon: Sword{Name: "长剑", Power: 20},
		Armor:  SteelArmor{Name: "钢甲", Protection: 10},
	}
	warrior.Fight("史莱姆")
	warrior.Defend()

	fmt.Println()

	// 示例2：匿名嵌入
	fmt.Println("【示例2：匿名嵌入】")
	hero := Hero{
		Name:       "英雄",
		Sword:      Sword{Name: "圣剑", Power: 50},
		SteelArmor: SteelArmor{Name: "神圣铠甲", Protection: 30},
	}
	// Hero 直接调用 Sword 的方法
	fmt.Printf("%s %s，伤害：%d\n", hero.Name, hero.Attack(), hero.Damage())
	fmt.Printf("%s %s，防御：%d\n", hero.Name, hero.Defend(), hero.Defense())

	fmt.Println()

	// 示例3：运行时动态更换装备（组合的核心优势）
	fmt.Println("【示例3：运行时动态更换装备】")
	archer := &Character{
		Name:   "弓箭手",
		Health: 80,
		Weapon: Bow{Name: "木弓", Power: 15},
	}
	archer.Fight("哥布林")

	// 战斗中更换武器！
	archer.EquipWeapon(Sword{Name: "短剑", Power: 25})
	archer.Fight("哥布林")

	fmt.Println()

	// 示例4：灵活组合不同能力
	fmt.Println("【示例4：灵活组合不同能力】")
	// 创建一个没有装备的角色
	novice := &Character{
		Name:   "新手",
		Health: 50,
	}
	novice.Fight("训练假人")
	novice.Defend()

	// 逐步装备
	novice.EquipWeapon(Bow{Name: "新手弓", Power: 10})
	novice.EquipArmor(SteelArmor{Name: "皮甲", Protection: 5})
	novice.Fight("训练假人")
	novice.Defend()
}
