/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/5/25 14:59
 * @File : hash.go
 */

package hash

import "novachat_engine/mtproto"

/*
public static int calcDocumentsHash(ArrayList<TLRPC.Document> arrayList, int maxCount) {
	if (arrayList == null) {
		return 0;
	}
	long acc = 0;
	for (int a = 0, N = Math.min(maxCount, arrayList.size()); a < N; a++) {
		TLRPC.Document document = arrayList.get(a);
		if (document == null) {
			continue;
		}
		int high_id = (int) (document.id >> 32);
		int lower_id = (int) document.id;
		acc = ((acc * 20261) + 0x80000000L + high_id) % 0x80000000L;
		acc = ((acc * 20261) + 0x80000000L + lower_id) % 0x80000000L;
	}
	return (int) acc;
}
*/

const MaxCountHash = 200
const UnLimitMaxCountHash = -1

func calcHash(acc int64, id int64) int64 {
	var high int64
	var low int64
	high = id >> 32
	low = id & 0xffffffff
	acc = ((acc * 20261) + 0x80000000 + high) % 0x80000000
	acc = ((acc * 20261) + 0x80000000 + low) % 0x80000000
	return acc
}

func DocumentsHash(documents []*mtproto.Document, maxCount int) int32 {
	if len(documents) == 0 {
		return 0
	}

	var (
		document *mtproto.Document
		acc      int64 = 0
	)
	for idx := 0; idx < len(documents); idx++ {
		if maxCount != UnLimitMaxCountHash {
			if idx >= maxCount {
				break
			}
		}

		document = documents[idx]
		if document == nil {
			continue
		}

		acc = calcHash(acc, document.Id)
	}

	return int32(acc)
}

func Hash(documents []int64, maxCount int) int32 {
	if len(documents) == 0 {
		return 0
	}

	var (
		acc int64 = 0
	)
	for idx := 0; idx < len(documents); idx++ {
		if UnLimitMaxCountHash != maxCount {
			if idx >= maxCount {
				break
			}
		}

		acc = calcHash(acc, documents[idx])
	}

	return int32(acc)
}

func Hash64(documents []int64, maxCount int) int64 {
	if len(documents) == 0 {
		return 0
	}

	var (
		acc int64 = 0
	)
	for idx := 0; idx < len(documents); idx++ {
		if UnLimitMaxCountHash != maxCount {
			if idx >= maxCount {
				break
			}
		}

		acc = calcHash(acc, documents[idx])
	}

	return acc
}

func SetHash(sets []*mtproto.StickerSet) int32 {
	var already uint32 = 0
	for _, v := range sets {
		if v.Id == 0 {
			continue
		}
		if v.Archived == true {
			continue
		}
		already = (already * 20261) + uint32(v.Hash)
	}

	return int32(already)
}
