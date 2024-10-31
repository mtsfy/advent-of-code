import hashlib
import multiprocessing


def find_hash(input, start, end, results):
    for i in range(start, end):
        input_value = f"{input}{i}"
        hash_result = hashlib.md5(input_value.encode('utf-8')).hexdigest()
        if hash_result.startswith("000000"):
            results.append(i)
            return
    

if __name__ == "__main__":
    input = "yzbqklnj"
    num_processes = multiprocessing.cpu_count()
    chunk_size = 1000000
    processes = []
    manager = multiprocessing.Manager()
    results = manager.list()  

    for p in range(num_processes):
        start = p * chunk_size + 1
        end = (p + 1) * chunk_size + 1
        process = multiprocessing.Process(target=find_hash, args=(input, start, end, results))
        processes.append(process)
        process.start()

    for process in processes:
        process.join()

    valid_hashes = [result for result in results if result is not None]
    
    if valid_hashes:
        print(f"The lowest number is: {min(valid_hashes)}")
