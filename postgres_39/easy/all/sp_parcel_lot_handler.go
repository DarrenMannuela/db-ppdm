package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcelLot(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel_lot")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel_lot

	for rows.Next() {
		var sp_parcel_lot dto.Sp_parcel_lot
		if err := rows.Scan(&sp_parcel_lot.Parcel_lot_id, &sp_parcel_lot.Parcel_lot_type, &sp_parcel_lot.Parcel_lot_num, &sp_parcel_lot.Active_ind, &sp_parcel_lot.Effective_date, &sp_parcel_lot.Expiry_date, &sp_parcel_lot.Gross_size, &sp_parcel_lot.Gross_size_ouom, &sp_parcel_lot.Lot_description, &sp_parcel_lot.Lot_name, &sp_parcel_lot.Parcel_congress_id, &sp_parcel_lot.Parcel_ohio_id, &sp_parcel_lot.Parcel_pbl_id, &sp_parcel_lot.Parcel_texas_id, &sp_parcel_lot.Ppdm_guid, &sp_parcel_lot.Remark, &sp_parcel_lot.Remark_type, &sp_parcel_lot.Source, &sp_parcel_lot.Spatial_description_id, &sp_parcel_lot.Spatial_obs_no, &sp_parcel_lot.Row_changed_by, &sp_parcel_lot.Row_changed_date, &sp_parcel_lot.Row_created_by, &sp_parcel_lot.Row_created_date, &sp_parcel_lot.Row_effective_date, &sp_parcel_lot.Row_expiry_date, &sp_parcel_lot.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel_lot to result
		result = append(result, sp_parcel_lot)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcelLot(c *fiber.Ctx) error {
	var sp_parcel_lot dto.Sp_parcel_lot

	setDefaults(&sp_parcel_lot)

	if err := json.Unmarshal(c.Body(), &sp_parcel_lot); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel_lot values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	sp_parcel_lot.Row_created_date = formatDate(sp_parcel_lot.Row_created_date)
	sp_parcel_lot.Row_changed_date = formatDate(sp_parcel_lot.Row_changed_date)
	sp_parcel_lot.Effective_date = formatDateString(sp_parcel_lot.Effective_date)
	sp_parcel_lot.Expiry_date = formatDateString(sp_parcel_lot.Expiry_date)
	sp_parcel_lot.Row_effective_date = formatDateString(sp_parcel_lot.Row_effective_date)
	sp_parcel_lot.Row_expiry_date = formatDateString(sp_parcel_lot.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_lot.Parcel_lot_id, sp_parcel_lot.Parcel_lot_type, sp_parcel_lot.Parcel_lot_num, sp_parcel_lot.Active_ind, sp_parcel_lot.Effective_date, sp_parcel_lot.Expiry_date, sp_parcel_lot.Gross_size, sp_parcel_lot.Gross_size_ouom, sp_parcel_lot.Lot_description, sp_parcel_lot.Lot_name, sp_parcel_lot.Parcel_congress_id, sp_parcel_lot.Parcel_ohio_id, sp_parcel_lot.Parcel_pbl_id, sp_parcel_lot.Parcel_texas_id, sp_parcel_lot.Ppdm_guid, sp_parcel_lot.Remark, sp_parcel_lot.Remark_type, sp_parcel_lot.Source, sp_parcel_lot.Spatial_description_id, sp_parcel_lot.Spatial_obs_no, sp_parcel_lot.Row_changed_by, sp_parcel_lot.Row_changed_date, sp_parcel_lot.Row_created_by, sp_parcel_lot.Row_created_date, sp_parcel_lot.Row_effective_date, sp_parcel_lot.Row_expiry_date, sp_parcel_lot.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcelLot(c *fiber.Ctx) error {
	var sp_parcel_lot dto.Sp_parcel_lot

	setDefaults(&sp_parcel_lot)

	if err := json.Unmarshal(c.Body(), &sp_parcel_lot); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel_lot.Parcel_lot_id = id

    if sp_parcel_lot.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel_lot set parcel_lot_type = :1, parcel_lot_num = :2, active_ind = :3, effective_date = :4, expiry_date = :5, gross_size = :6, gross_size_ouom = :7, lot_description = :8, lot_name = :9, parcel_congress_id = :10, parcel_ohio_id = :11, parcel_pbl_id = :12, parcel_texas_id = :13, ppdm_guid = :14, remark = :15, remark_type = :16, source = :17, spatial_description_id = :18, spatial_obs_no = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where parcel_lot_id = :27")
	if err != nil {
		return err
	}

	sp_parcel_lot.Row_changed_date = formatDate(sp_parcel_lot.Row_changed_date)
	sp_parcel_lot.Effective_date = formatDateString(sp_parcel_lot.Effective_date)
	sp_parcel_lot.Expiry_date = formatDateString(sp_parcel_lot.Expiry_date)
	sp_parcel_lot.Row_effective_date = formatDateString(sp_parcel_lot.Row_effective_date)
	sp_parcel_lot.Row_expiry_date = formatDateString(sp_parcel_lot.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_lot.Parcel_lot_type, sp_parcel_lot.Parcel_lot_num, sp_parcel_lot.Active_ind, sp_parcel_lot.Effective_date, sp_parcel_lot.Expiry_date, sp_parcel_lot.Gross_size, sp_parcel_lot.Gross_size_ouom, sp_parcel_lot.Lot_description, sp_parcel_lot.Lot_name, sp_parcel_lot.Parcel_congress_id, sp_parcel_lot.Parcel_ohio_id, sp_parcel_lot.Parcel_pbl_id, sp_parcel_lot.Parcel_texas_id, sp_parcel_lot.Ppdm_guid, sp_parcel_lot.Remark, sp_parcel_lot.Remark_type, sp_parcel_lot.Source, sp_parcel_lot.Spatial_description_id, sp_parcel_lot.Spatial_obs_no, sp_parcel_lot.Row_changed_by, sp_parcel_lot.Row_changed_date, sp_parcel_lot.Row_created_by, sp_parcel_lot.Row_effective_date, sp_parcel_lot.Row_expiry_date, sp_parcel_lot.Row_quality, sp_parcel_lot.Parcel_lot_id)
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

func PatchSpParcelLot(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel_lot set "
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
	query += " where parcel_lot_id = :id"

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

func DeleteSpParcelLot(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel_lot dto.Sp_parcel_lot
	sp_parcel_lot.Parcel_lot_id = id

	stmt, err := db.Prepare("delete from sp_parcel_lot where parcel_lot_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel_lot.Parcel_lot_id)
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


