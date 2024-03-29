/*
	Created On:6 Feb 2017
	Author:Nilav

	Implementing Test Cases for testing different user management operations
 */
package test_test

import (
	"github.com/structure/futsalStruct"
	"../programs/service"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"encoding/json"
	"fmt"
)

var _ = Describe("UserTest", func() {

	var (
		header futsalStruct.Header
		user futsalStruct.User
		jsonData string
	)

	BeforeEach(func() {
		header=futsalStruct.Header{"Add User","ADD"}
		user=futsalStruct.User{"1","Nilav","something@gmail.com","Tech","Khatiwada","password","9843241699"}

		b,_:=json.Marshal(user)
		jsonData=string(b)
	})

	Describe("Testing the User Activity", func() {
		Context("Testing the User Signup",func() {

			It("Should return successfully signed up",func() {
				fmt.Println("There is something   ",jsonData)
				Expect(service.AddUser(user)).To(Equal("Successfull"))
			})

		})

		Context("Testing User Login",func() {

			It("Should return successfully signed up",func() {
				fmt.Println("There is something   ",jsonData)
				Expect(service.LoginUser(user)).To(Equal("Successfully Logged In"))
			})

		})

		Context("Getting User By User Id",func() {

			It("Should return successfully signed up",func() {
				fmt.Println("There is something   ",jsonData)
				Expect(service.GetUserById(user.UserId)).To(Equal("Successfull"))
			})

		Context("Updating User",func() {

			It("Should return successfully signed up",func() {
				fmt.Println("There is something   ",jsonData)
				Expect(service.UpdateUserById(user.UserId,user)).To(Equal("Successfull"))
			})

			})

		Context("Deleting User",func() {

			It("Should return successfully signed up",func() {
				fmt.Println("There is something   ",jsonData)
				Expect(service.DeleteUserById(user.UserId)).To(Equal("Successfull"))
			})

			})
		})
	})

	Describe("Login", func() {
		Context("Testing the User Login", func() {

			It("Should return successfully logged in", func() {
				fmt.Println("There is something   ", jsonData)
				Expect(service.LoginUser(user)).To(Equal("Successfull"))
			})

		})
	})
})
