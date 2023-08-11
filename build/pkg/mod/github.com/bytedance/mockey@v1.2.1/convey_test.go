/*
 * Copyright 2022 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mockey

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

//go:noinline
func Fun0() string {
	// TODO: deal with branch optimization
	ok1 := Fun1()
	fmt.Printf("Fun0: call Fun1, %v\n", ok1)
	if !ok1 {
		return "exit"
	}
	ok2 := Fun2()
	fmt.Printf("Fun0: call Fun2, %v\n", ok2)
	if ok2 {
		return "fun2"
	}
	ok3 := Fun3()
	fmt.Printf("Fun0: call Fun3, %v\n", ok3)
	if ok3 {
		return "fun3"
	}
	return "xxx"
}

//go:noinline
func Fun1() bool {
	fmt.Println("Fun1")
	return false
}

//go:noinline
func Fun2() bool {
	fmt.Println("Fun2")
	return false
}

//go:noinline
func Fun3() bool {
	fmt.Println("Fun3")
	return false
}

func TestConvey(t *testing.T) {
	PatchConvey("test", t, func() {
		Mock(Fun1).Return(true).Build()

		PatchConvey("test case 2", func() {
			m2 := Mock(Fun2).Return(true).Build()
			m3 := Mock(Fun3).Return(true).Build()

			r := Fun0()
			convey.So(r, convey.ShouldEqual, "fun2")
			convey.So(m2.Times(), convey.ShouldEqual, 1)
			convey.So(m3.Times(), convey.ShouldEqual, 0)
		})

		PatchConvey("test case 3", func() {
			m2 := Mock(Fun2).Build()
			m3 := Mock(Fun3).Return(true).Build()

			r := Fun0()
			convey.So(r, convey.ShouldEqual, "fun3")
			convey.So(m2.Times(), convey.ShouldEqual, 1)
			convey.So(m3.Times(), convey.ShouldEqual, 1)
		})

		PatchConvey("test case mock times", func() {
			m2 := Mock(Fun2).When(func() bool { return false }).Build()
			m3 := Mock(Fun3).Return(true).Build()

			r := Fun0()
			convey.So(r, convey.ShouldEqual, "fun3")
			convey.So(m2.Times(), convey.ShouldEqual, 1)
			convey.So(m2.MockTimes(), convey.ShouldEqual, 0)
			convey.So(m3.Times(), convey.ShouldEqual, 1)
		})
	})
}
