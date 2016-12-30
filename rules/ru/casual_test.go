package ru_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/ru"
)

func TestCasualDate(t *testing.T) {
	fixt := []Fixture{
		{"Это нужно сделать прямо сейчас", 44, "сейчас", 0},
		{"Это нужно сделать сегодня", 33, "сегодня", 0},
		{"Это нужно сделать завтра вечером", 33, "завтра", time.Hour * 24},
		{"Это нужно сделать вчера вечером", 33, "вчера", -(time.Hour * 24)},
	}

	w := when.New(nil)
	w.Add(ru.CasualDate(rules.Skip))

	ApplyFixtures(t, "ru.CasualDate", w, fixt)
}

func TestCasualTime(t *testing.T) {
	fixt := []Fixture{
		{"Это нужно сделать этим утром ", 33, "этим утром", 8 * time.Hour},
		{"Это нужно сделать до обеда", 33, "до обеда", 12 * time.Hour},
		{"Это нужно сделать после обеда", 33, "после обеда", 15 * time.Hour},
		{"Это нужно сделать к вечеру", 33, "к вечеру", 18 * time.Hour},
	}

	w := when.New(nil)
	w.Add(ru.CasualTime(rules.Skip))

	ApplyFixtures(t, "ru.CasualTime", w, fixt)
}

func TestCasualDateCasualTime(t *testing.T) {
	fixt := []Fixture{
		{"Это нужно сделать завтра после обеда", 33, "завтра после обеда", (15 + 24) * time.Hour},
	}

	w := when.New(nil)
	w.Add(
		ru.CasualDate(rules.Skip),
		ru.CasualTime(rules.OverWrite),
	)

	ApplyFixtures(t, "ru.CasualDate|ru.CasualTime", w, fixt)
}
