package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcel(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel

	for rows.Next() {
		var sp_parcel dto.Sp_parcel
		if err := rows.Scan(&sp_parcel.Spatial_description_id, &sp_parcel.Spatial_obs_no, &sp_parcel.Parcel_id, &sp_parcel.Active_ind, &sp_parcel.Description, &sp_parcel.Dls_road_allowance_id, &sp_parcel.Effective_date, &sp_parcel.Expiry_date, &sp_parcel.Gross_size, &sp_parcel.Gross_size_ouom, &sp_parcel.Inactivation_date, &sp_parcel.Local_coord_system_id, &sp_parcel.Parcel_carter_id, &sp_parcel.Parcel_congress_id, &sp_parcel.Parcel_dls_id, &sp_parcel.Parcel_fps_id, &sp_parcel.Parcel_libya_id, &sp_parcel.Parcel_lot_id, &sp_parcel.Parcel_lot_num, &sp_parcel.Parcel_lot_type, &sp_parcel.Parcel_ne_loc_id, &sp_parcel.Parcel_north_sea_id, &sp_parcel.Parcel_nts_id, &sp_parcel.Parcel_offshore_id, &sp_parcel.Parcel_ohio_id, &sp_parcel.Parcel_pbl_id, &sp_parcel.Parcel_texas_id, &sp_parcel.Percent_ownership, &sp_parcel.Ppdm_guid, &sp_parcel.Remark, &sp_parcel.Source, &sp_parcel.Row_changed_by, &sp_parcel.Row_changed_date, &sp_parcel.Row_created_by, &sp_parcel.Row_created_date, &sp_parcel.Row_effective_date, &sp_parcel.Row_expiry_date, &sp_parcel.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel to result
		result = append(result, sp_parcel)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcel(c *fiber.Ctx) error {
	var sp_parcel dto.Sp_parcel

	setDefaults(&sp_parcel)

	if err := json.Unmarshal(c.Body(), &sp_parcel); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38)")
	if err != nil {
		return err
	}
	sp_parcel.Row_created_date = formatDate(sp_parcel.Row_created_date)
	sp_parcel.Row_changed_date = formatDate(sp_parcel.Row_changed_date)
	sp_parcel.Effective_date = formatDateString(sp_parcel.Effective_date)
	sp_parcel.Expiry_date = formatDateString(sp_parcel.Expiry_date)
	sp_parcel.Inactivation_date = formatDateString(sp_parcel.Inactivation_date)
	sp_parcel.Row_effective_date = formatDateString(sp_parcel.Row_effective_date)
	sp_parcel.Row_expiry_date = formatDateString(sp_parcel.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel.Spatial_description_id, sp_parcel.Spatial_obs_no, sp_parcel.Parcel_id, sp_parcel.Active_ind, sp_parcel.Description, sp_parcel.Dls_road_allowance_id, sp_parcel.Effective_date, sp_parcel.Expiry_date, sp_parcel.Gross_size, sp_parcel.Gross_size_ouom, sp_parcel.Inactivation_date, sp_parcel.Local_coord_system_id, sp_parcel.Parcel_carter_id, sp_parcel.Parcel_congress_id, sp_parcel.Parcel_dls_id, sp_parcel.Parcel_fps_id, sp_parcel.Parcel_libya_id, sp_parcel.Parcel_lot_id, sp_parcel.Parcel_lot_num, sp_parcel.Parcel_lot_type, sp_parcel.Parcel_ne_loc_id, sp_parcel.Parcel_north_sea_id, sp_parcel.Parcel_nts_id, sp_parcel.Parcel_offshore_id, sp_parcel.Parcel_ohio_id, sp_parcel.Parcel_pbl_id, sp_parcel.Parcel_texas_id, sp_parcel.Percent_ownership, sp_parcel.Ppdm_guid, sp_parcel.Remark, sp_parcel.Source, sp_parcel.Row_changed_by, sp_parcel.Row_changed_date, sp_parcel.Row_created_by, sp_parcel.Row_created_date, sp_parcel.Row_effective_date, sp_parcel.Row_expiry_date, sp_parcel.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcel(c *fiber.Ctx) error {
	var sp_parcel dto.Sp_parcel

	setDefaults(&sp_parcel)

	if err := json.Unmarshal(c.Body(), &sp_parcel); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel.Spatial_description_id = id

    if sp_parcel.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel set spatial_obs_no = :1, parcel_id = :2, active_ind = :3, description = :4, dls_road_allowance_id = :5, effective_date = :6, expiry_date = :7, gross_size = :8, gross_size_ouom = :9, inactivation_date = :10, local_coord_system_id = :11, parcel_carter_id = :12, parcel_congress_id = :13, parcel_dls_id = :14, parcel_fps_id = :15, parcel_libya_id = :16, parcel_lot_id = :17, parcel_lot_num = :18, parcel_lot_type = :19, parcel_ne_loc_id = :20, parcel_north_sea_id = :21, parcel_nts_id = :22, parcel_offshore_id = :23, parcel_ohio_id = :24, parcel_pbl_id = :25, parcel_texas_id = :26, percent_ownership = :27, ppdm_guid = :28, remark = :29, source = :30, row_changed_by = :31, row_changed_date = :32, row_created_by = :33, row_effective_date = :34, row_expiry_date = :35, row_quality = :36 where spatial_description_id = :38")
	if err != nil {
		return err
	}

	sp_parcel.Row_changed_date = formatDate(sp_parcel.Row_changed_date)
	sp_parcel.Effective_date = formatDateString(sp_parcel.Effective_date)
	sp_parcel.Expiry_date = formatDateString(sp_parcel.Expiry_date)
	sp_parcel.Inactivation_date = formatDateString(sp_parcel.Inactivation_date)
	sp_parcel.Row_effective_date = formatDateString(sp_parcel.Row_effective_date)
	sp_parcel.Row_expiry_date = formatDateString(sp_parcel.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel.Spatial_obs_no, sp_parcel.Parcel_id, sp_parcel.Active_ind, sp_parcel.Description, sp_parcel.Dls_road_allowance_id, sp_parcel.Effective_date, sp_parcel.Expiry_date, sp_parcel.Gross_size, sp_parcel.Gross_size_ouom, sp_parcel.Inactivation_date, sp_parcel.Local_coord_system_id, sp_parcel.Parcel_carter_id, sp_parcel.Parcel_congress_id, sp_parcel.Parcel_dls_id, sp_parcel.Parcel_fps_id, sp_parcel.Parcel_libya_id, sp_parcel.Parcel_lot_id, sp_parcel.Parcel_lot_num, sp_parcel.Parcel_lot_type, sp_parcel.Parcel_ne_loc_id, sp_parcel.Parcel_north_sea_id, sp_parcel.Parcel_nts_id, sp_parcel.Parcel_offshore_id, sp_parcel.Parcel_ohio_id, sp_parcel.Parcel_pbl_id, sp_parcel.Parcel_texas_id, sp_parcel.Percent_ownership, sp_parcel.Ppdm_guid, sp_parcel.Remark, sp_parcel.Source, sp_parcel.Row_changed_by, sp_parcel.Row_changed_date, sp_parcel.Row_created_by, sp_parcel.Row_effective_date, sp_parcel.Row_expiry_date, sp_parcel.Row_quality, sp_parcel.Spatial_description_id)
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

func PatchSpParcel(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "inactivation_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where spatial_description_id = :id"

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

func DeleteSpParcel(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel dto.Sp_parcel
	sp_parcel.Spatial_description_id = id

	stmt, err := db.Prepare("delete from sp_parcel where spatial_description_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel.Spatial_description_id)
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


