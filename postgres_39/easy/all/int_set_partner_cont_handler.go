package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetIntSetPartnerCont(c *fiber.Ctx) error {
	rows, err := db.Query("select * from int_set_partner_cont")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Int_set_partner_cont

	for rows.Next() {
		var int_set_partner_cont dto.Int_set_partner_cont
		if err := rows.Scan(&int_set_partner_cont.Interest_set_id, &int_set_partner_cont.Interest_set_seq_no, &int_set_partner_cont.Partner_ba_id, &int_set_partner_cont.Partner_obs_no, &int_set_partner_cont.Contact_ba_id, &int_set_partner_cont.Contact_obs_no, &int_set_partner_cont.Contact_id, &int_set_partner_cont.Active_ind, &int_set_partner_cont.Address_obs_no, &int_set_partner_cont.Address_source, &int_set_partner_cont.Contact_org_ba_id, &int_set_partner_cont.Contact_org_ba_seq_no, &int_set_partner_cont.Contact_org_organization_id, &int_set_partner_cont.Contact_role, &int_set_partner_cont.Contract_id, &int_set_partner_cont.Effective_date, &int_set_partner_cont.Expiry_date, &int_set_partner_cont.Inactive_date, &int_set_partner_cont.Ppdm_guid, &int_set_partner_cont.Provision_id, &int_set_partner_cont.Remark, &int_set_partner_cont.Source, &int_set_partner_cont.Row_changed_by, &int_set_partner_cont.Row_changed_date, &int_set_partner_cont.Row_created_by, &int_set_partner_cont.Row_created_date, &int_set_partner_cont.Row_effective_date, &int_set_partner_cont.Row_expiry_date, &int_set_partner_cont.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append int_set_partner_cont to result
		result = append(result, int_set_partner_cont)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetIntSetPartnerCont(c *fiber.Ctx) error {
	var int_set_partner_cont dto.Int_set_partner_cont

	setDefaults(&int_set_partner_cont)

	if err := json.Unmarshal(c.Body(), &int_set_partner_cont); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into int_set_partner_cont values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	int_set_partner_cont.Row_created_date = formatDate(int_set_partner_cont.Row_created_date)
	int_set_partner_cont.Row_changed_date = formatDate(int_set_partner_cont.Row_changed_date)
	int_set_partner_cont.Effective_date = formatDateString(int_set_partner_cont.Effective_date)
	int_set_partner_cont.Expiry_date = formatDateString(int_set_partner_cont.Expiry_date)
	int_set_partner_cont.Inactive_date = formatDateString(int_set_partner_cont.Inactive_date)
	int_set_partner_cont.Row_effective_date = formatDateString(int_set_partner_cont.Row_effective_date)
	int_set_partner_cont.Row_expiry_date = formatDateString(int_set_partner_cont.Row_expiry_date)






	rows, err := stmt.Exec(int_set_partner_cont.Interest_set_id, int_set_partner_cont.Interest_set_seq_no, int_set_partner_cont.Partner_ba_id, int_set_partner_cont.Partner_obs_no, int_set_partner_cont.Contact_ba_id, int_set_partner_cont.Contact_obs_no, int_set_partner_cont.Contact_id, int_set_partner_cont.Active_ind, int_set_partner_cont.Address_obs_no, int_set_partner_cont.Address_source, int_set_partner_cont.Contact_org_ba_id, int_set_partner_cont.Contact_org_ba_seq_no, int_set_partner_cont.Contact_org_organization_id, int_set_partner_cont.Contact_role, int_set_partner_cont.Contract_id, int_set_partner_cont.Effective_date, int_set_partner_cont.Expiry_date, int_set_partner_cont.Inactive_date, int_set_partner_cont.Ppdm_guid, int_set_partner_cont.Provision_id, int_set_partner_cont.Remark, int_set_partner_cont.Source, int_set_partner_cont.Row_changed_by, int_set_partner_cont.Row_changed_date, int_set_partner_cont.Row_created_by, int_set_partner_cont.Row_created_date, int_set_partner_cont.Row_effective_date, int_set_partner_cont.Row_expiry_date, int_set_partner_cont.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateIntSetPartnerCont(c *fiber.Ctx) error {
	var int_set_partner_cont dto.Int_set_partner_cont

	setDefaults(&int_set_partner_cont)

	if err := json.Unmarshal(c.Body(), &int_set_partner_cont); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	int_set_partner_cont.Interest_set_id = id

    if int_set_partner_cont.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update int_set_partner_cont set interest_set_seq_no = :1, partner_ba_id = :2, partner_obs_no = :3, contact_ba_id = :4, contact_obs_no = :5, contact_id = :6, active_ind = :7, address_obs_no = :8, address_source = :9, contact_org_ba_id = :10, contact_org_ba_seq_no = :11, contact_org_organization_id = :12, contact_role = :13, contract_id = :14, effective_date = :15, expiry_date = :16, inactive_date = :17, ppdm_guid = :18, provision_id = :19, remark = :20, source = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where interest_set_id = :29")
	if err != nil {
		return err
	}

	int_set_partner_cont.Row_changed_date = formatDate(int_set_partner_cont.Row_changed_date)
	int_set_partner_cont.Effective_date = formatDateString(int_set_partner_cont.Effective_date)
	int_set_partner_cont.Expiry_date = formatDateString(int_set_partner_cont.Expiry_date)
	int_set_partner_cont.Inactive_date = formatDateString(int_set_partner_cont.Inactive_date)
	int_set_partner_cont.Row_effective_date = formatDateString(int_set_partner_cont.Row_effective_date)
	int_set_partner_cont.Row_expiry_date = formatDateString(int_set_partner_cont.Row_expiry_date)






	rows, err := stmt.Exec(int_set_partner_cont.Interest_set_seq_no, int_set_partner_cont.Partner_ba_id, int_set_partner_cont.Partner_obs_no, int_set_partner_cont.Contact_ba_id, int_set_partner_cont.Contact_obs_no, int_set_partner_cont.Contact_id, int_set_partner_cont.Active_ind, int_set_partner_cont.Address_obs_no, int_set_partner_cont.Address_source, int_set_partner_cont.Contact_org_ba_id, int_set_partner_cont.Contact_org_ba_seq_no, int_set_partner_cont.Contact_org_organization_id, int_set_partner_cont.Contact_role, int_set_partner_cont.Contract_id, int_set_partner_cont.Effective_date, int_set_partner_cont.Expiry_date, int_set_partner_cont.Inactive_date, int_set_partner_cont.Ppdm_guid, int_set_partner_cont.Provision_id, int_set_partner_cont.Remark, int_set_partner_cont.Source, int_set_partner_cont.Row_changed_by, int_set_partner_cont.Row_changed_date, int_set_partner_cont.Row_created_by, int_set_partner_cont.Row_effective_date, int_set_partner_cont.Row_expiry_date, int_set_partner_cont.Row_quality, int_set_partner_cont.Interest_set_id)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchIntSetPartnerCont(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update int_set_partner_cont set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "effective_date" || key == "expiry_date" || key == "inactive_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where interest_set_id = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeleteIntSetPartnerCont(c *fiber.Ctx) error {
	id := c.Params("id")
	var int_set_partner_cont dto.Int_set_partner_cont
	int_set_partner_cont.Interest_set_id = id

	stmt, err := db.Prepare("delete from int_set_partner_cont where interest_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(int_set_partner_cont.Interest_set_id)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}


