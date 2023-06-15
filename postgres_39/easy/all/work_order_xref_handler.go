package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWorkOrderXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from work_order_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Work_order_xref

	for rows.Next() {
		var work_order_xref dto.Work_order_xref
		if err := rows.Scan(&work_order_xref.Work_order_id, &work_order_xref.Work_order_xref_id, &work_order_xref.Active_ind, &work_order_xref.Business_associate_id, &work_order_xref.Effective_date, &work_order_xref.Expiry_date, &work_order_xref.Parent_work_order_id, &work_order_xref.Ppdm_guid, &work_order_xref.Reference_id, &work_order_xref.Remark, &work_order_xref.Source, &work_order_xref.Wo_xref_type, &work_order_xref.Row_changed_by, &work_order_xref.Row_changed_date, &work_order_xref.Row_created_by, &work_order_xref.Row_created_date, &work_order_xref.Row_effective_date, &work_order_xref.Row_expiry_date, &work_order_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append work_order_xref to result
		result = append(result, work_order_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWorkOrderXref(c *fiber.Ctx) error {
	var work_order_xref dto.Work_order_xref

	setDefaults(&work_order_xref)

	if err := json.Unmarshal(c.Body(), &work_order_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into work_order_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	work_order_xref.Row_created_date = formatDate(work_order_xref.Row_created_date)
	work_order_xref.Row_changed_date = formatDate(work_order_xref.Row_changed_date)
	work_order_xref.Effective_date = formatDateString(work_order_xref.Effective_date)
	work_order_xref.Expiry_date = formatDateString(work_order_xref.Expiry_date)
	work_order_xref.Row_effective_date = formatDateString(work_order_xref.Row_effective_date)
	work_order_xref.Row_expiry_date = formatDateString(work_order_xref.Row_expiry_date)






	rows, err := stmt.Exec(work_order_xref.Work_order_id, work_order_xref.Work_order_xref_id, work_order_xref.Active_ind, work_order_xref.Business_associate_id, work_order_xref.Effective_date, work_order_xref.Expiry_date, work_order_xref.Parent_work_order_id, work_order_xref.Ppdm_guid, work_order_xref.Reference_id, work_order_xref.Remark, work_order_xref.Source, work_order_xref.Wo_xref_type, work_order_xref.Row_changed_by, work_order_xref.Row_changed_date, work_order_xref.Row_created_by, work_order_xref.Row_created_date, work_order_xref.Row_effective_date, work_order_xref.Row_expiry_date, work_order_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWorkOrderXref(c *fiber.Ctx) error {
	var work_order_xref dto.Work_order_xref

	setDefaults(&work_order_xref)

	if err := json.Unmarshal(c.Body(), &work_order_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	work_order_xref.Work_order_id = id

    if work_order_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update work_order_xref set work_order_xref_id = :1, active_ind = :2, business_associate_id = :3, effective_date = :4, expiry_date = :5, parent_work_order_id = :6, ppdm_guid = :7, reference_id = :8, remark = :9, source = :10, wo_xref_type = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where work_order_id = :19")
	if err != nil {
		return err
	}

	work_order_xref.Row_changed_date = formatDate(work_order_xref.Row_changed_date)
	work_order_xref.Effective_date = formatDateString(work_order_xref.Effective_date)
	work_order_xref.Expiry_date = formatDateString(work_order_xref.Expiry_date)
	work_order_xref.Row_effective_date = formatDateString(work_order_xref.Row_effective_date)
	work_order_xref.Row_expiry_date = formatDateString(work_order_xref.Row_expiry_date)






	rows, err := stmt.Exec(work_order_xref.Work_order_xref_id, work_order_xref.Active_ind, work_order_xref.Business_associate_id, work_order_xref.Effective_date, work_order_xref.Expiry_date, work_order_xref.Parent_work_order_id, work_order_xref.Ppdm_guid, work_order_xref.Reference_id, work_order_xref.Remark, work_order_xref.Source, work_order_xref.Wo_xref_type, work_order_xref.Row_changed_by, work_order_xref.Row_changed_date, work_order_xref.Row_created_by, work_order_xref.Row_effective_date, work_order_xref.Row_expiry_date, work_order_xref.Row_quality, work_order_xref.Work_order_id)
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

func PatchWorkOrderXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update work_order_xref set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where work_order_id = :id"

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

func DeleteWorkOrderXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var work_order_xref dto.Work_order_xref
	work_order_xref.Work_order_id = id

	stmt, err := db.Prepare("delete from work_order_xref where work_order_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(work_order_xref.Work_order_id)
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


