package main

import (
    "bytes"
    "compress/gzip"
    "fmt"
    "io"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func gatherers_ab_dbi() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xca, 0xcc,
		0x2b, 0x2e, 0x49, 0xcc, 0xc9, 0x49, 0x4d, 0x51, 0x50, 0x4a, 0x29, 0xca,
		0x2f, 0x48, 0x4a, 0x4d, 0x2c, 0x52, 0x28, 0x2e, 0xce, 0x50, 0x48, 0xce,
		0xc9, 0x4c, 0xcd, 0x2b, 0x51, 0x80, 0x4b, 0x2b, 0x29, 0xc0, 0xa4, 0xb9,
		0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x97, 0x69, 0xa5, 0x33, 0x00,
		0x00, 0x00,
		},
		"gatherers/ab/dbi",
	)
}

func gatherers_ab_gen() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xb4, 0x55,
		0x6d, 0x6f, 0xe2, 0x46, 0x10, 0xfe, 0x9e, 0x5f, 0x31, 0xb2, 0xb8, 0xda,
		0x56, 0x89, 0x81, 0xbc, 0x9d, 0x94, 0xe8, 0xa4, 0x46, 0x21, 0xd7, 0xd0,
		0x4b, 0x8b, 0x15, 0x9c, 0xf0, 0xe1, 0x38, 0x71, 0x8b, 0xbd, 0xc6, 0x5b,
		0xec, 0xdd, 0xed, 0xee, 0x9a, 0x0b, 0xd7, 0xde, 0x7f, 0xef, 0xac, 0x8d,
		0x21, 0x90, 0x97, 0x6f, 0x27, 0x25, 0xe0, 0x99, 0xe7, 0x79, 0x66, 0x3c,
		0x33, 0x3b, 0x4b, 0x5a, 0xf2, 0xd8, 0x30, 0xc1, 0x21, 0x15, 0x25, 0x4f,
		0x3c, 0x1f, 0xfe, 0x3d, 0x00, 0x18, 0xde, 0x47, 0x1f, 0xbe, 0xd2, 0xb9,
		0xa2, 0x12, 0x0e, 0x73, 0xa6, 0xc0, 0xf1, 0x5a, 0x3d, 0xdf, 0x81, 0x4e,
		0x26, 0x0a, 0x0a, 0xff, 0xfd, 0x02, 0x35, 0xb2, 0x04, 0x27, 0xa4, 0xaa,
		0x60, 0x5a, 0x5b, 0x7d, 0x42, 0x39, 0xa3, 0x89, 0xf3, 0x15, 0xe5, 0xa2,
		0x34, 0xb2, 0x34, 0x53, 0xa2, 0x14, 0x59, 0x81, 0xd3, 0x3a, 0x72, 0xf0,
		0x03, 0x43, 0x3a, 0x07, 0x3f, 0x0e, 0x0e, 0xaa, 0x34, 0xe0, 0x6a, 0x8a,
		0xc9, 0x62, 0x45, 0xbe, 0xaf, 0xda, 0xd0, 0x6d, 0x43, 0x4b, 0xb3, 0xef,
		0xb4, 0x0d, 0x5a, 0xc4, 0x0b, 0x92, 0x24, 0x6a, 0xca, 0xb8, 0xd7, 0x92,
		0x42, 0x19, 0x44, 0x98, 0x75, 0xf8, 0xfe, 0x85, 0x0b, 0x6e, 0x92, 0x08,
		0xed, 0x36, 0x11, 0xfa, 0x7d, 0x31, 0x62, 0x7c, 0x0e, 0x26, 0xa3, 0x30,
		0x08, 0xf7, 0xd1, 0x68, 0x38, 0xbc, 0x85, 0x7e, 0x7f, 0x38, 0x82, 0xfe,
		0xe0, 0xfa, 0x99, 0x94, 0xea, 0x45, 0x52, 0xce, 0xb4, 0xa1, 0x72, 0x1f,
		0x1a, 0x2b, 0x66, 0x0c, 0xe5, 0x30, 0x5b, 0xc1, 0x48, 0x98, 0x04, 0x61,
		0x29, 0xb0, 0xbe, 0x59, 0x4e, 0xc1, 0xf2, 0xc0, 0xb3, 0x5e, 0xdf, 0xdd,
		0xd4, 0x11, 0xae, 0x4c, 0x86, 0xc5, 0x3f, 0x94, 0x39, 0xa7, 0x8a, 0xcc,
		0x58, 0xce, 0x0c, 0x2a, 0x63, 0xc2, 0xd1, 0xb4, 0xe2, 0x1a, 0x5e, 0xee,
		0xc0, 0x7a, 0x0d, 0x37, 0x31, 0x2c, 0x5d, 0x63, 0xff, 0x15, 0x7c, 0x8c,
		0x42, 0xd0, 0x54, 0x2d, 0xa9, 0xd2, 0x28, 0x4e, 0x8d, 0xac, 0xb8, 0xdb,
		0x6c, 0x37, 0xf7, 0xb7, 0x9f, 0xa0, 0xff, 0x3b, 0x44, 0x84, 0x43, 0x2c,
		0xf8, 0x1c, 0x49, 0xba, 0xd4, 0x92, 0xc5, 0x4c, 0x94, 0x1a, 0xb9, 0x8a,
		0x49, 0xb3, 0x09, 0x1b, 0x46, 0x27, 0xef, 0xe1, 0xd2, 0x18, 0x12, 0x2f,
		0x60, 0x64, 0x88, 0x32, 0x34, 0x79, 0x99, 0xdf, 0x08, 0x3e, 0x92, 0x98,
		0xce, 0x84, 0x58, 0xc0, 0x95, 0x42, 0x0d, 0x55, 0xf0, 0x80, 0xef, 0x81,
		0xa3, 0xb5, 0xaf, 0xd2, 0x40, 0x64, 0x56, 0x6a, 0xba, 0x49, 0xb1, 0x71,
		0xcf, 0x54, 0x69, 0x28, 0x56, 0x10, 0x57, 0x55, 0xbf, 0xc2, 0x5e, 0x95,
		0x49, 0x46, 0x82, 0x79, 0xc9, 0x75, 0x8e, 0x83, 0xa3, 0xea, 0xb7, 0x79,
		0x41, 0x58, 0x1e, 0xc4, 0xa2, 0x78, 0xda, 0xe6, 0x5d, 0x31, 0x78, 0x95,
		0xca, 0xdf, 0x04, 0xc9, 0x8c, 0x91, 0xe7, 0x9d, 0x0e, 0x29, 0x8d, 0xc8,
		0xd9, 0x82, 0xa6, 0xb3, 0x40, 0xa8, 0x79, 0x27, 0x3d, 0x0d, 0x32, 0x53,
		0xe4, 0x3b, 0xb9, 0xd7, 0x8c, 0x8d, 0x12, 0x1d, 0x99, 0x50, 0xe7, 0xf0,
		0x69, 0x10, 0xc1, 0xcd, 0xf5, 0xdd, 0xf0, 0xf9, 0x6c, 0x1b, 0x64, 0x9b,
		0x6d, 0x94, 0x8b, 0x6f, 0xb9, 0x50, 0xcc, 0x8e, 0x43, 0x37, 0xcf, 0xb0,
		0x73, 0x62, 0xfe, 0xa4, 0x86, 0x68, 0x99, 0x0b, 0x66, 0x90, 0x53, 0x6c,
		0x8d, 0x06, 0x2f, 0x18, 0xa7, 0x78, 0xd0, 0x53, 0x33, 0xad, 0x27, 0xeb,
		0x3e, 0x71, 0xad, 0x87, 0xfd, 0x94, 0x8b, 0x9d, 0x99, 0x4a, 0x25, 0x1e,
		0x57, 0x6e, 0x63, 0x42, 0x6d, 0xee, 0x6c, 0x4e, 0x9a, 0x0b, 0x91, 0xb4,
		0x41, 0xe2, 0x9c, 0x3c, 0x87, 0xb4, 0xa4, 0x5d, 0x1f, 0xa7, 0xed, 0x54,
		0x6e, 0xc7, 0xaf, 0x56, 0xca, 0x62, 0xd3, 0xdd, 0x85, 0xda, 0xdb, 0xa8,
		0x1f, 0xfb, 0x87, 0xdf, 0x88, 0x05, 0x1e, 0xfd, 0x0f, 0xa0, 0x08, 0x4f,
		0x44, 0x11, 0xc4, 0x99, 0x60, 0x31, 0xf5, 0x3e, 0x4f, 0xdc, 0xe0, 0xaf,
		0xeb, 0x08, 0xae, 0x6e, 0xef, 0x26, 0x6e, 0x1b, 0x26, 0xee, 0xe8, 0xa1,
		0x57, 0x3f, 0x44, 0x04, 0x7b, 0x67, 0x20, 0xbc, 0xaa, 0xcd, 0x31, 0xe3,
		0x67, 0x27, 0x17, 0x30, 0xb8, 0x3c, 0x3b, 0xd9, 0x71, 0x3c, 0x6e, 0xec,
		0xe1, 0xd8, 0x3e, 0x7e, 0xf1, 0xe1, 0x57, 0xb4, 0x2e, 0xf0, 0x7f, 0xff,
		0x0d, 0x7e, 0x62, 0x71, 0x8a, 0xfe, 0x53, 0x52, 0x6d, 0xb0, 0x3c, 0x27,
		0x1c, 0x8e, 0x22, 0x78, 0xa7, 0x3b, 0x8f, 0x45, 0xae, 0x64, 0x1c, 0xc8,
		0x4c, 0xc2, 0x4d, 0x14, 0x85, 0x9d, 0x5e, 0xd0, 0x9d, 0xa8, 0x09, 0xbf,
		0x11, 0xda, 0x9c, 0x23, 0x6e, 0x9f, 0xef, 0x71, 0x42, 0x87, 0x97, 0x73,
		0xca, 0xd1, 0x33, 0xe0, 0x86, 0x2a, 0x4e, 0x72, 0x18, 0x0b, 0x95, 0x48,
		0x45, 0xb5, 0x86, 0xbb, 0xf0, 0xca, 0x2e, 0x22, 0x8e, 0xd3, 0xd4, 0x8b,
		0xb2, 0x93, 0xd2, 0x5b, 0x37, 0xd2, 0x7e, 0x31, 0x6e, 0xbc, 0xb3, 0x36,
		0x1c, 0x1f, 0x61, 0x01, 0x7b, 0xee, 0x5e, 0x17, 0x2b, 0x3a, 0xea, 0x76,
		0xbb, 0xcf, 0x21, 0x04, 0x7a, 0x2f, 0xf9, 0x4f, 0x8f, 0x4f, 0xdb, 0x70,
		0x7a, 0xfc, 0xfe, 0x85, 0x60, 0x98, 0xe3, 0xcc, 0xf7, 0xab, 0xeb, 0x84,
		0x01, 0xe3, 0x16, 0x9f, 0xd3, 0xca, 0xdd, 0xf5, 0xe1, 0x0b, 0x76, 0xfe,
		0x33, 0xb8, 0xeb, 0xc5, 0xee, 0xbc, 0x4b, 0x82, 0xea, 0x0f, 0x46, 0x24,
		0x25, 0x8a, 0xa1, 0xbd, 0x5f, 0x42, 0x41, 0xb4, 0xbd, 0x82, 0x0e, 0x89,
		0x94, 0x41, 0xe6, 0x6e, 0x6c, 0x48, 0x59, 0xbe, 0xdd, 0xac, 0x3f, 0x84,
		0x28, 0x72, 0xf2, 0xea, 0x25, 0xf8, 0x77, 0x0d, 0xbf, 0x72, 0x09, 0x36,
		0x41, 0xaa, 0xb4, 0xe0, 0x32, 0xbe, 0x24, 0x9a, 0x2d, 0x29, 0x2c, 0xe8,
		0xea, 0x1b, 0xf6, 0xb9, 0xfe, 0x5d, 0x02, 0xcf, 0xc2, 0xdb, 0xbd, 0x24,
		0xd5, 0xb5, 0xf6, 0x06, 0xbd, 0x26, 0x6c, 0x05, 0xd5, 0xe1, 0x79, 0x83,
		0x5f, 0xe1, 0x48, 0xff, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x56, 0xff, 0xa7,
		0xdb, 0x0f, 0x07, 0x00, 0x00,
		},
		"gatherers/ab/gen",
	)
}

func gatherers_ab_hii() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xca, 0xcc,
		0x2b, 0x2e, 0x49, 0xcc, 0xc9, 0x49, 0x4d, 0x51, 0x50, 0xca, 0x28, 0xc8,
		0xcc, 0x4b, 0x37, 0x56, 0x80, 0x8b, 0x28, 0x29, 0x40, 0x44, 0xb8, 0x00,
		0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0xa6, 0xa6, 0xc9, 0x24, 0x00, 0x00,
		0x00,
		},
		"gatherers/ab/hii",
	)
}

func gatherers_ab_psi() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xca, 0xcc,
		0x2b, 0x2e, 0x49, 0xcc, 0xc9, 0x49, 0x4d, 0x51, 0x50, 0xcf, 0xcb, 0x4d,
		0x2c, 0x50, 0x28, 0xc8, 0x2f, 0x2a, 0x29, 0x4e, 0x4e, 0xcc, 0xcb, 0x4b,
		0x2d, 0x52, 0x80, 0xcb, 0xa9, 0x43, 0x24, 0xd5, 0xb9, 0x90, 0x55, 0x27,
		0x27, 0x96, 0xe0, 0x56, 0x0d, 0x94, 0x44, 0x51, 0x5d, 0x90, 0x07, 0x52,
		0x86, 0x53, 0x3d, 0x44, 0x1a, 0x45, 0x47, 0x76, 0x5e, 0x7e, 0x72, 0x36,
		0x50, 0x19, 0x2e, 0x2d, 0x50, 0x79, 0x54, 0x37, 0x25, 0x16, 0xe3, 0x71,
		0x13, 0x50, 0x52, 0x9d, 0x0b, 0x10, 0x00, 0x00, 0xff, 0xff, 0x30, 0xf1,
		0xb0, 0xe7, 0xf0, 0x00, 0x00, 0x00,
		},
		"gatherers/ab/psi",
	)
}

func gatherers_ab_wbd() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8c, 0x8c,
		0x41, 0x0e, 0xc2, 0x30, 0x0c, 0x04, 0xef, 0xbc, 0xc2, 0xf2, 0x07, 0xfc,
		0x24, 0x44, 0xa8, 0x53, 0x2c, 0x15, 0xb6, 0xb2, 0x1b, 0x25, 0xcf, 0x27,
		0x22, 0xe4, 0x80, 0xc4, 0xa1, 0x97, 0x9d, 0xbd, 0xcc, 0x2c, 0xe6, 0x57,
		0x6d, 0x16, 0x47, 0x10, 0x0b, 0xf6, 0x43, 0x56, 0x60, 0xdd, 0x54, 0xee,
		0x0f, 0xc7, 0x73, 0x82, 0x89, 0xc7, 0xa1, 0xaa, 0x89, 0x92, 0xa3, 0x86,
		0x3a, 0x65, 0x94, 0xd7, 0xc2, 0x97, 0x9f, 0x42, 0x09, 0x97, 0xcd, 0x92,
		0x64, 0x73, 0xcd, 0x68, 0x93, 0x3d, 0xf0, 0x7d, 0xa7, 0x0b, 0xd8, 0xd5,
		0x6f, 0x63, 0xbb, 0xfd, 0xe1, 0x3f, 0xf7, 0x1d, 0x00, 0x00, 0xff, 0xff,
		0x86, 0xed, 0xbd, 0x13, 0xc0, 0x00, 0x00, 0x00,
		},
		"gatherers/ab/wbd",
	)
}

func gatherers_an_bac() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xca, 0x2b,
		0xcd, 0x4d, 0x4a, 0x2d, 0x8a, 0xcf, 0x4f, 0x8b, 0xcf, 0xc9, 0xcc, 0x4b,
		0x2d, 0x56, 0x50, 0xd2, 0x4b, 0x4a, 0x2c, 0xce, 0x28, 0x4a, 0x56, 0xe2,
		0x02, 0x04, 0x00, 0x00, 0xff, 0xff, 0x42, 0xbd, 0x37, 0xd7, 0x1a, 0x00,
		0x00, 0x00,
		},
		"gatherers/an/bac",
	)
}

func gatherers_an_crd() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x72, 0x0e,
		0xf2, 0xf7, 0x8b, 0xf7, 0x0f, 0x0d, 0xb1, 0x55, 0xd1, 0x48, 0x2e, 0xca,
		0xcf, 0x2b, 0x49, 0x4c, 0x52, 0xd0, 0xcd, 0x51, 0xa8, 0x51, 0x48, 0x2f,
		0x4a, 0x2d, 0x50, 0xd0, 0x2d, 0x53, 0x50, 0x52, 0x56, 0x02, 0xf2, 0x8a,
		0x53, 0x53, 0x14, 0xd4, 0xf5, 0xe3, 0x62, 0x8a, 0xb5, 0x54, 0xf4, 0x53,
		0xd4, 0x35, 0xb9, 0xf2, 0x4b, 0x4b, 0x0a, 0x4a, 0x4b, 0xe2, 0x13, 0x8b,
		0x8a, 0x12, 0x2b, 0x15, 0x94, 0x40, 0x1a, 0x15, 0x52, 0xf3, 0x4a, 0x8a,
		0x32, 0x53, 0x8b, 0x15, 0x52, 0x52, 0x4b, 0x52, 0x93, 0x4b, 0x52, 0x53,
		0x94, 0x14, 0x94, 0x54, 0xaa, 0x61, 0xa6, 0xeb, 0xeb, 0xab, 0x6b, 0xa9,
		0xd7, 0x2a, 0x71, 0x01, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xe0, 0x2a,
		0xd9, 0x6e, 0x00, 0x00, 0x00,
		},
		"gatherers/an/crd",
	)
}

func gatherers_an_doi() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xca, 0xcc,
		0x2b, 0x2e, 0x49, 0xcc, 0xc9, 0x49, 0x4d, 0x51, 0x50, 0x4a, 0xc9, 0x4f,
		0xce, 0x4e, 0x2d, 0x52, 0x80, 0x8b, 0x28, 0x29, 0x40, 0x44, 0xb8, 0x00,
		0x01, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x6e, 0x3b, 0x32, 0x24, 0x00, 0x00,
		0x00,
		},
		"gatherers/an/doi",
	)
}

func gatherers_an_esi() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x4a, 0xc9,
		0x2c, 0x8a, 0x4f, 0xad, 0xc8, 0x2c, 0x2e, 0x29, 0x56, 0x50, 0xd2, 0x2f,
		0x2d, 0x2e, 0xd2, 0x2f, 0xce, 0x48, 0x2c, 0x4a, 0xd5, 0x4f, 0xcd, 0x49,
		0x2c, 0x2e, 0xc9, 0x4c, 0x2e, 0x4e, 0x4d, 0x2c, 0x4a, 0xce, 0x50, 0x52,
		0x50, 0x42, 0xe1, 0x2b, 0xa4, 0xe5, 0x97, 0xe6, 0xa5, 0x28, 0x71, 0x01,
		0x02, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x10, 0xaf, 0x69, 0x3c, 0x00, 0x00,
		0x00,
		},
		"gatherers/an/esi",
	)
}

func gatherers_an_fic() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xca, 0x2f,
		0x2d, 0x29, 0x28, 0x2d, 0x89, 0x2f, 0xcd, 0xcb, 0x49, 0x2d, 0x2e, 0x8e,
		0x4f, 0x4b, 0xcc, 0x29, 0xae, 0x54, 0x50, 0x4a, 0xcb, 0x2c, 0xce, 0x50,
		0xc8, 0xc9, 0xcc, 0x4b, 0x55, 0x48, 0xce, 0x2f, 0xcd, 0x2b, 0x51, 0x52,
		0x50, 0xf1, 0x0b, 0xf5, 0x75, 0x72, 0x0d, 0x52, 0x50, 0xd1, 0xc8, 0x2b,
		0xcd, 0x4d, 0x4a, 0x2d, 0x8a, 0xcf, 0x4f, 0x8b, 0x07, 0xc9, 0x17, 0x2b,
		0xd4, 0xe9, 0xeb, 0x25, 0xe7, 0xe7, 0xa5, 0x65, 0xa6, 0xeb, 0x83, 0x34,
		0xe9, 0x43, 0xd8, 0x7a, 0x20, 0xb6, 0x26, 0x17, 0x20, 0x00, 0x00, 0xff,
		0xff, 0xf4, 0x17, 0xf6, 0x40, 0x5c, 0x00, 0x00, 0x00,
		},
		"gatherers/an/fic",
	)
}

func gatherers_an_mef() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x0c, 0xcb,
		0x4d, 0x0a, 0xc2, 0x30, 0x10, 0x05, 0xe0, 0xbd, 0xa7, 0x78, 0x14, 0x31,
		0xad, 0x50, 0x5d, 0xe8, 0x41, 0xdc, 0x8b, 0x48, 0x30, 0xaf, 0x3a, 0x60,
		0x7e, 0x98, 0x4c, 0x05, 0x21, 0x87, 0x6f, 0x36, 0xdf, 0xee, 0xcb, 0xab,
		0x95, 0xd5, 0x9e, 0x5e, 0xd5, 0xff, 0xe1, 0x22, 0x83, 0x78, 0x2c, 0xf2,
		0xe5, 0x58, 0x27, 0x04, 0x1a, 0x5f, 0xc6, 0xe0, 0xb0, 0x1f, 0x17, 0x49,
		0x01, 0xe7, 0x4f, 0x8e, 0xc4, 0x9c, 0x7c, 0x77, 0x38, 0x9e, 0x62, 0xb9,
		0x5f, 0xda, 0xf5, 0x31, 0xa0, 0x1d, 0xf0, 0x56, 0x16, 0xcc, 0x3f, 0xb8,
		0x1b, 0x35, 0x4a, 0xad, 0x92, 0x53, 0xff, 0x49, 0xfa, 0x9e, 0x76, 0x5b,
		0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0xb9, 0xe1, 0x87, 0x65, 0x00, 0x00,
		0x00,
		},
		"gatherers/an/mef",
	)
}

func gatherers_an_moi() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xca, 0xcc,
		0x2b, 0x2e, 0x49, 0xcc, 0xc9, 0x49, 0x4d, 0x51, 0x50, 0xca, 0xcd, 0xcf,
		0x4b, 0xcf, 0x4f, 0x49, 0x52, 0x80, 0x0b, 0x29, 0x29, 0x40, 0x84, 0xb8,
		0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0x94, 0xde, 0xa3, 0x94, 0x25, 0x00,
		0x00, 0x00,
		},
		"gatherers/an/moi",
	)
}

func gatherers_an_myi() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xca, 0xcc,
		0x2b, 0x2e, 0x49, 0xcc, 0xc9, 0x49, 0x4d, 0x51, 0x50, 0xca, 0xad, 0x2c,
		0x2e, 0xcc, 0x51, 0x80, 0x0b, 0x28, 0x29, 0x80, 0x05, 0xb8, 0x00, 0x01,
		0x00, 0x00, 0xff, 0xff, 0xf3, 0xbc, 0xfd, 0xf3, 0x22, 0x00, 0x00, 0x00,
		},
		"gatherers/an/myi",
	)
}

func gatherers_an_nus() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x34, 0x8e,
		0xcd, 0x4a, 0x86, 0x40, 0x18, 0x85, 0xd7, 0xf3, 0x5e, 0xc5, 0xe1, 0x63,
		0x16, 0x4a, 0x88, 0x75, 0x01, 0x6e, 0xfa, 0x59, 0x08, 0x65, 0x60, 0xba,
		0x12, 0x11, 0xd3, 0x77, 0x72, 0x60, 0x9a, 0x91, 0xf9, 0x29, 0x82, 0x2e,
		0x3e, 0x33, 0xda, 0x9c, 0xc5, 0xe1, 0x3c, 0x0f, 0xe7, 0xb1, 0x7e, 0xe9,
		0xa6, 0xfb, 0xba, 0xad, 0x64, 0x66, 0x02, 0xca, 0xcd, 0xbd, 0x33, 0xbe,
		0xf1, 0xe6, 0x79, 0x47, 0xf1, 0x81, 0xf4, 0x9a, 0x6c, 0x4c, 0x39, 0xd1,
		0xdd, 0x73, 0xdf, 0x74, 0xd5, 0x35, 0x29, 0xe7, 0xb1, 0x6a, 0x0f, 0x6d,
		0x21, 0xff, 0x51, 0x5a, 0x1d, 0x09, 0xad, 0x30, 0x0c, 0x28, 0x8e, 0x3a,
		0x3b, 0x61, 0xf9, 0xbb, 0x2a, 0x39, 0x2e, 0xe5, 0x3e, 0x87, 0xf0, 0xb9,
		0xe6, 0x18, 0x47, 0x12, 0x71, 0x63, 0x4b, 0x42, 0x18, 0x8e, 0xf8, 0x33,
		0x9e, 0x79, 0x75, 0x43, 0x42, 0xe9, 0x43, 0x63, 0x99, 0x5c, 0x8a, 0x7b,
		0x8a, 0x53, 0xb2, 0x86, 0x43, 0x98, 0xd4, 0x6c, 0xc2, 0x17, 0x2e, 0x29,
		0xb0, 0xc7, 0xe2, 0x8e, 0x2f, 0x17, 0xc8, 0xa6, 0x7f, 0xba, 0x7d, 0x68,
		0x21, 0x4f, 0x94, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x5b, 0x02,
		0x53, 0xc0, 0x00, 0x00, 0x00,
		},
		"gatherers/an/nus",
	)
}

func gatherers_an_poi() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xca, 0xcc,
		0x2b, 0x2e, 0x49, 0xcc, 0xc9, 0x49, 0x4d, 0x51, 0x50, 0x2a, 0xc8, 0x2f,
		0x2e, 0x49, 0x2f, 0x4a, 0x2d, 0x2e, 0xcc, 0x51, 0x80, 0x8b, 0x2a, 0x29,
		0x14, 0x00, 0xf9, 0x5c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0xbd,
		0xc0, 0xb3, 0x26, 0x00, 0x00, 0x00,
		},
		"gatherers/an/poi",
	)
}

func gatherers_an_sii() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xca, 0xcc,
		0x2b, 0x2e, 0x49, 0xcc, 0xc9, 0x49, 0x4d, 0x51, 0x50, 0x2a, 0x2e, 0xcc,
		0xc9, 0x2c, 0x49, 0x55, 0x80, 0x8b, 0x28, 0x29, 0x40, 0x44, 0xb8, 0x00,
		0x01, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xf9, 0x50, 0x9f, 0x24, 0x00, 0x00,
		0x00,
		},
		"gatherers/an/sii",
	)
}

func gatherers_an_trn() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x2c, 0xcb,
		0x41, 0x0a, 0xc2, 0x40, 0x0c, 0x85, 0xe1, 0xbd, 0xa7, 0x78, 0x0c, 0x62,
		0x5b, 0xa1, 0x7a, 0x15, 0x6f, 0x20, 0x83, 0x79, 0xd5, 0x80, 0x4d, 0x4a,
		0x26, 0x15, 0x04, 0x0f, 0x6f, 0x17, 0xdd, 0xfc, 0xab, 0xff, 0xf3, 0x35,
		0x97, 0x35, 0xef, 0x35, 0xa2, 0x7e, 0xd1, 0xa5, 0x47, 0xd0, 0x12, 0x93,
		0xbe, 0xd9, 0xb7, 0x01, 0xc2, 0xe4, 0x23, 0x29, 0x1d, 0x8e, 0xfd, 0xa4,
		0x26, 0xb8, 0xbe, 0x7c, 0x26, 0x46, 0xab, 0x5b, 0xcb, 0xf9, 0xb2, 0xff,
		0x05, 0xbf, 0x13, 0x9e, 0xc1, 0x05, 0xe3, 0x07, 0xe5, 0xc6, 0x98, 0xb5,
		0x35, 0x75, 0xdb, 0xbc, 0x29, 0xa5, 0x0c, 0x87, 0x7f, 0x00, 0x00, 0x00,
		0xff, 0xff, 0x2e, 0x82, 0xae, 0x48, 0x67, 0x00, 0x00, 0x00,
		},
		"gatherers/an/trn",
	)
}

func gatherers_an_vrc() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xca, 0x2b,
		0xcd, 0x4d, 0x4a, 0x2d, 0x8a, 0xcf, 0x4f, 0x8b, 0xcf, 0xc9, 0xcc, 0x4b,
		0x2d, 0x56, 0x50, 0xd2, 0x2b, 0xcb, 0xcc, 0x2d, 0x4a, 0x56, 0xe2, 0x02,
		0x04, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x67, 0x80, 0x6b, 0x19, 0x00, 0x00,
		0x00,
		},
		"gatherers/an/vrc",
	)
}

func gatherers_an_zsc() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xca, 0x2b,
		0xcd, 0x4d, 0x4a, 0x2d, 0x8a, 0xcf, 0x4f, 0x8b, 0xcf, 0xc9, 0xcc, 0x4b,
		0x2d, 0x56, 0x50, 0xd2, 0xab, 0x2a, 0xce, 0x28, 0x4a, 0x56, 0xe2, 0x02,
		0x04, 0x00, 0x00, 0xff, 0xff, 0x65, 0xa3, 0x8d, 0xc6, 0x19, 0x00, 0x00,
		0x00,
		},
		"gatherers/an/zsc",
	)
}

func gatherers_common() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x84, 0x54,
		0x4d, 0x6f, 0xdb, 0x30, 0x0c, 0xbd, 0xfb, 0x57, 0x70, 0x84, 0x50, 0xb7,
		0x41, 0x1d, 0x2f, 0x09, 0x76, 0x59, 0xe1, 0x01, 0xed, 0xd0, 0xdb, 0xda,
		0x0e, 0x05, 0x76, 0x18, 0xea, 0xc0, 0x70, 0x2d, 0xb9, 0x11, 0x26, 0x4b,
		0x9e, 0x25, 0xf7, 0x63, 0x49, 0xf6, 0xdb, 0x47, 0x39, 0xce, 0x92, 0xcc,
		0x01, 0x76, 0x91, 0x98, 0x27, 0xea, 0x91, 0xef, 0x89, 0xf1, 0xd5, 0xdd,
		0xdd, 0x97, 0xeb, 0xcb, 0xdb, 0x04, 0x1f, 0x8d, 0x51, 0x22, 0xd7, 0x18,
		0xdc, 0x7e, 0xbb, 0xb9, 0xba, 0xbe, 0x4f, 0x50, 0xb7, 0xd5, 0xa3, 0x68,
		0x30, 0xb8, 0xbc, 0xbf, 0xbf, 0xfc, 0x9e, 0x60, 0xde, 0x34, 0xf9, 0x1b,
		0x06, 0x41, 0xd9, 0xea, 0xc2, 0x49, 0xa3, 0xc1, 0xb4, 0xae, 0x6e, 0xdd,
		0xe9, 0x19, 0x2c, 0x03, 0x00, 0x2b, 0x1c, 0x44, 0x25, 0x05, 0xa2, 0x58,
		0x18, 0xc0, 0x65, 0x8a, 0x3a, 0xaf, 0x44, 0x8a, 0x1f, 0x53, 0x64, 0x93,
		0x14, 0xcf, 0x21, 0x45, 0xf7, 0x56, 0xf7, 0xc0, 0x74, 0x03, 0x3c, 0xe7,
		0xaa, 0xf5, 0x08, 0x9b, 0xad, 0x31, 0x58, 0x0f, 0x98, 0xb3, 0xae, 0x62,
		0xcf, 0x2f, 0x4b, 0x78, 0x78, 0x80, 0x48, 0x03, 0x9b, 0xc2, 0x7c, 0x4e,
		0x80, 0x5b, 0x08, 0x4d, 0xdb, 0x5e, 0x65, 0x80, 0xcf, 0x77, 0x37, 0x57,
		0x09, 0x3e, 0x60, 0xf7, 0xa3, 0x34, 0x0d, 0x48, 0x90, 0xfe, 0xc6, 0x05,
		0x70, 0xd3, 0x61, 0xdb, 0x14, 0xe6, 0x37, 0x6a, 0x44, 0xa6, 0xb8, 0x49,
		0xe6, 0x46, 0x8b, 0x1d, 0x45, 0x77, 0x8c, 0x73, 0xdc, 0x43, 0x96, 0x7e,
		0x8b, 0xe3, 0x14, 0x53, 0xa4, 0xe5, 0x3c, 0xc5, 0x75, 0x77, 0xb8, 0xe9,
		0x14, 0x48, 0x23, 0xd2, 0xd2, 0x39, 0xe5, 0x83, 0xee, 0x3e, 0x25, 0x94,
		0xf2, 0x98, 0xae, 0x56, 0x2b, 0x61, 0x6d, 0x56, 0xe6, 0xca, 0x1e, 0xca,
		0x63, 0x33, 0x52, 0x28, 0xe0, 0x3d, 0x29, 0x84, 0x93, 0x93, 0x1e, 0x79,
		0x97, 0x00, 0xfa, 0x54, 0x81, 0x3b, 0xd8, 0xfb, 0x30, 0xfb, 0xc7, 0x87,
		0xc3, 0x56, 0xa6, 0x7e, 0x99, 0x1d, 0xe9, 0x61, 0xf3, 0xaa, 0x99, 0x29,
		0x33, 0x25, 0xb5, 0xb0, 0x7d, 0x7d, 0x2e, 0x9b, 0x84, 0x4d, 0x83, 0x9d,
		0xd1, 0xbf, 0x86, 0x46, 0xfb, 0x1c, 0xec, 0x19, 0xbb, 0x55, 0x09, 0xff,
		0xc6, 0x09, 0x9b, 0x6c, 0x08, 0x6c, 0xc2, 0x4e, 0x4b, 0xa9, 0x39, 0xc4,
		0x0b, 0x53, 0x89, 0x78, 0x14, 0x33, 0x02, 0x21, 0xaa, 0x08, 0x12, 0xb5,
		0x5b, 0xc0, 0x84, 0xe2, 0xfc, 0xf5, 0x6f, 0xec, 0xef, 0x02, 0xdb, 0xb2,
		0xc0, 0xea, 0x04, 0x9e, 0x1a, 0x51, 0x43, 0xf4, 0x0c, 0xf8, 0x55, 0x34,
		0x95, 0xb4, 0xd6, 0xb7, 0xcb, 0x85, 0x96, 0x82, 0xe3, 0x99, 0xaf, 0x78,
		0xc4, 0x3e, 0x2f, 0x17, 0xbc, 0x10, 0x28, 0x4c, 0xab, 0x1d, 0x02, 0xdb,
		0x4c, 0x2f, 0xb0, 0xd3, 0x97, 0x02, 0x22, 0x05, 0xbe, 0x07, 0x0b, 0x2b,
		0x70, 0xb9, 0x54, 0x10, 0x4d, 0x28, 0xca, 0x5f, 0x7e, 0x40, 0xb8, 0xac,
		0x1b, 0xa9, 0x1d, 0xb0, 0xc9, 0x3a, 0x3c, 0x3b, 0xb0, 0x47, 0x49, 0xeb,
		0x32, 0xdf, 0x54, 0xe6, 0xc7, 0x75, 0x6b, 0x4f, 0xc1, 0xe1, 0xf7, 0x05,
		0x74, 0xe2, 0xc6, 0x7b, 0x2a, 0xa6, 0x10, 0xf9, 0x2c, 0x28, 0x7b, 0x35,
		0xe1, 0x68, 0x3c, 0x0a, 0x29, 0x36, 0x34, 0x91, 0x75, 0x4e, 0x09, 0x38,
		0x8a, 0xc7, 0x4f, 0xa2, 0x8a, 0x47, 0x78, 0x88, 0x8e, 0x75, 0x3d, 0x04,
		0xb5, 0xe1, 0x22, 0xab, 0x0c, 0x6f, 0x49, 0x9e, 0x3f, 0x5c, 0xd1, 0x68,
		0x73, 0x08, 0xed, 0x6a, 0x3c, 0x4a, 0xc7, 0xab, 0x55, 0xe8, 0x01, 0xd3,
		0x38, 0xda, 0x5a, 0x2d, 0x7f, 0x42, 0x54, 0x6c, 0x81, 0x48, 0x0f, 0x25,
		0x3c, 0x49, 0x97, 0x91, 0x9d, 0xe6, 0x3f, 0x0a, 0x3e, 0xf4, 0x0a, 0xf8,
		0x4e, 0x01, 0xdd, 0x0c, 0x0f, 0xf8, 0xa4, 0xb6, 0x2e, 0x57, 0x4a, 0xf0,
		0xdd, 0xac, 0xbe, 0x2c, 0x64, 0xb1, 0xf0, 0x13, 0xf2, 0x09, 0x62, 0x2e,
		0x9e, 0x63, 0xdd, 0x2a, 0x75, 0x71, 0x7c, 0x14, 0x59, 0xff, 0x81, 0x01,
		0xd7, 0xb4, 0x62, 0x38, 0x8e, 0xf4, 0x3c, 0x99, 0x78, 0xa5, 0x8e, 0xed,
		0xde, 0x3f, 0x01, 0x22, 0x4e, 0x4f, 0x03, 0xf3, 0x21, 0x25, 0x95, 0x3c,
		0x4a, 0xf8, 0x27, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xee, 0xbc, 0xb3, 0xc6,
		0x04, 0x00, 0x00,
		},
		"gatherers/common",
	)
}


// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string] func() ([]byte, error) {
	"gatherers/ab/dbi": gatherers_ab_dbi,
	"gatherers/ab/gen": gatherers_ab_gen,
	"gatherers/ab/hii": gatherers_ab_hii,
	"gatherers/ab/psi": gatherers_ab_psi,
	"gatherers/ab/wbd": gatherers_ab_wbd,
	"gatherers/an/bac": gatherers_an_bac,
	"gatherers/an/crd": gatherers_an_crd,
	"gatherers/an/doi": gatherers_an_doi,
	"gatherers/an/esi": gatherers_an_esi,
	"gatherers/an/fic": gatherers_an_fic,
	"gatherers/an/mef": gatherers_an_mef,
	"gatherers/an/moi": gatherers_an_moi,
	"gatherers/an/myi": gatherers_an_myi,
	"gatherers/an/nus": gatherers_an_nus,
	"gatherers/an/poi": gatherers_an_poi,
	"gatherers/an/sii": gatherers_an_sii,
	"gatherers/an/trn": gatherers_an_trn,
	"gatherers/an/vrc": gatherers_an_vrc,
	"gatherers/an/zsc": gatherers_an_zsc,
	"gatherers/common": gatherers_common,

}
