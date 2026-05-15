import { ref, onMounted } from 'vue'
import axios from 'axios'

export function useDropdownSalary() {
    const salaryData = ref(null)
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/home/dropdown/salary')
            salaryData.value = res.data
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { salaryData, loading, error }
}
