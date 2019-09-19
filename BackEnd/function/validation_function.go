/*created by H.Mlk*/
package function
import(
	_"github.com/kataras/iris"
	_"github.com/kataras/iris/context"
	//"v2/AbsharAutomation/v2.1/config"
	///"strings"
	_"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


func IsValidId(id string) bool{
	return id=="" || !bson.IsObjectIdHex(id)

}

func IPValidate(ip string) bool{ //admin Ips will be added later
	//if ip ==config.IpAccessBranch
	if ip=="11112222"{// 32 digit ip address of admin user
		return true
	}

	return false

}


/*

	///////////////


	// Gets a single user
	// Method:   GET
	// Resource: this to get a user by msisdn
	app.Handle("GET", "/users/{username: string}", func(ctx context.Context) {
		username := ctx.Params().Get("username")
		//fmt.Println(msisdn)
		//if msisdn == "" {
			//ctx.JSON(context.Map{"response": "please pass a valid username"})
		//}


		//var user model.User////???????????????/
		err := r.db.C(r.collection).Find(bson.M{"first_name":FirstName})//////is it and??

		if err != nil {
			//return nil ,err
			return err
		}

		//return &users , nil
		//return &users////?????????
		return nil   ////////////////???????????????????????????????????????????????



		result := User{}
		err = c.Find(bson.M{"username": username}).One(&result)
		if err != nil {
			ctx.JSON(context.Map{"response": err.Error()})
		}
		ctx.JSON(context.Map{"response": result})
	})

}

*/

