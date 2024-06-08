
total_path_list=[[1257, 1261, 1262, 1263], [1257, 1261], [1256, 1259, 1262, 1263], [1256, 1259]]
final_path = []
for path in total_path_list:
    find_index = None
    # Merge if path take same node
    for ind, target_path in enumerate(final_path):
        if set(path[1:]) & set(target_path[1:]) and path[-1] == target_path[-1]:
            find_index = ind
            break
    if find_index is not None:
        target_path = final_path[find_index]
        # Merge, except for sink nodes
        final_path[find_index] = sorted(set(target_path + path[:-1]))
    else:
        final_path.append(path)

print(final_path)